package did

import (
	"encoding/json"
	"fmt"
	"net/url"
	"regexp"
	"strings"
)

// DID is parsed according to the generic syntax: https://w3c.github.io/did-core/#generic-did-syntax
type DID struct {
	Scheme           string `json:"scheme"`             // Scheme is always "did"
	Method           string `json:"method"`             // Method is the specific DID methods
	MethodSpecificID string `json:"method_specific_id"` // MethodSpecificID is the unique ID computed or assigned by the DID method
}

// String returns a string representation of this DID.
func (d *DID) String() string {
	return fmt.Sprintf("%s:%s:%s", d.Scheme, d.Method, d.MethodSpecificID)
}

// Parse parses the string according to the generic DID syntax.
// See https://w3c.github.io/did-core/#generic-did-syntax.
func Parse(did string) (*DID, error) {
	// I could not find a good ABNF parser :(
	const idChar = `a-zA-Z0-9-_\.`
	const methodChar = `a-zA-Z0-9`
	regex := fmt.Sprintf(`^did:[a-z]+:(:+|[:%s]+)*[%%:%s]+[^:]$`, methodChar, idChar)

	r, err := regexp.Compile(regex)
	if err != nil {
		return nil, fmt.Errorf("failed to compile regex=%s (this should not have happened!). %w", regex, err)
	}

	if !r.MatchString(did) {
		return nil, fmt.Errorf(
			"invalid did: %s. Make sure it conforms to the DID syntax: https://w3c.github.io/did-core/#did-syntax", did)
	}

	parts := strings.SplitN(did, ":", 3)

	return &DID{
		Scheme:           "did",
		Method:           parts[1],
		MethodSpecificID: parts[2],
	}, nil
}

func (d *DID) Marshal() ([]byte, error) {
	return json.Marshal(d)
}

// DIDURL holds a DID URL.
type DIDURL struct { // nolint:golint // ignore name stutter
	DID
	Path     string
	Queries  map[string][]string
	Fragment string
}

// ParseDIDURL parses a DID URL string into a DIDURL object.
func ParseDIDURL(didURL string) (*DIDURL, error) {
	split := strings.IndexAny(didURL, "?/#")

	didPart := didURL
	pathQueryFragment := ""

	if split != -1 {
		didPart = didURL[:split]
		pathQueryFragment = didURL[split:]
	}

	retDID, err := Parse(didPart)
	if err != nil {
		return nil, err
	}

	if pathQueryFragment == "" {
		return &DIDURL{
			DID:     *retDID,
			Queries: map[string][]string{},
		}, nil
	}

	hasPath := pathQueryFragment[0] == '/'

	if !hasPath {
		pathQueryFragment = "/" + pathQueryFragment
	}

	urlParts, err := url.Parse(pathQueryFragment)
	if err != nil {
		return nil, fmt.Errorf("failed to parse path, query, and fragment components of DID URL: %w", err)
	}

	ret := &DIDURL{
		DID:      *retDID,
		Queries:  urlParts.Query(),
		Fragment: urlParts.Fragment,
	}

	if hasPath {
		ret.Path = urlParts.Path
	}

	return ret, nil
}
