package did

import (
	"fmt"
	"path/filepath"
	"reflect"
	"runtime"
	"testing"
)

func TestIsURL(t *testing.T) {
	t.Run("returns false if no Path or Fragment", func(t *testing.T) {
		d := &DID{Method: "example", ID: "123"}
		Assert(t, false, d.IsURL())
	})

	t.Run("returns true if Params", func(t *testing.T) {
		d := &DID{Method: "example", ID: "123", Params: []Param{{Name: "foo", Value: "bar"}}}
		Assert(t, true, d.IsURL())
	})

	t.Run("returns true if Path", func(t *testing.T) {
		d := &DID{Method: "example", ID: "123", Path: "a/b"}
		Assert(t, true, d.IsURL())
	})

	t.Run("returns true if PathSegements", func(t *testing.T) {
		d := &DID{Method: "example", ID: "123", PathSegments: []string{"a", "b"}}
		Assert(t, true, d.IsURL())
	})

	t.Run("returns true if Query", func(t *testing.T) {
		d := &DID{Method: "example", ID: "123", Query: "abc"}
		Assert(t, true, d.IsURL())
	})

	t.Run("returns true if Fragment", func(t *testing.T) {
		d := &DID{Method: "example", ID: "123", Fragment: "00000"}
		Assert(t, true, d.IsURL())
	})

	t.Run("returns true if Path and Fragment", func(t *testing.T) {
		d := &DID{Method: "example", ID: "123", Path: "a/b", Fragment: "00000"}
		Assert(t, true, d.IsURL())
	})
}

func TestString(t *testing.T) {
	t.Run("assembles a DID", func(t *testing.T) {
		d := &DID{Method: "example", ID: "123"}
		Assert(t, "did:example:123", d.String())
	})

	t.Run("assembles a DID from IDStrings", func(t *testing.T) {
		d := &DID{Method: "example", IDStrings: []string{"123", "456"}}
		Assert(t, "did:example:123:456", d.String())
	})

	t.Run("returns empty string if no method", func(t *testing.T) {
		d := &DID{ID: "123"}
		Assert(t, "", d.String())
	})

	t.Run("returns empty string in no ID or IDStrings", func(t *testing.T) {
		d := &DID{Method: "example"}
		Assert(t, "", d.String())
	})

	t.Run("returns empty string if Param Name does not exist", func(t *testing.T) {
		d := &DID{Method: "example", ID: "123", Params: []Param{{Name: "", Value: "agent"}}}
		Assert(t, "", d.String())
	})

	t.Run("returns name string if Param Value does not exist", func(t *testing.T) {
		d := &DID{Method: "example", ID: "123", Params: []Param{{Name: "service", Value: ""}}}
		Assert(t, "did:example:123;service", d.String())
	})

	t.Run("returns param string with name and value", func(t *testing.T) {
		d := &DID{Method: "example", ID: "123", Params: []Param{{Name: "service", Value: "agent"}}}
		Assert(t, "did:example:123;service=agent", d.String())
	})

	t.Run("includes Param generic", func(t *testing.T) {
		d := &DID{Method: "example", ID: "123", Params: []Param{{Name: "service", Value: "agent"}}}
		Assert(t, "did:example:123;service=agent", d.String())
	})

	t.Run("includes Param method", func(t *testing.T) {
		d := &DID{Method: "example", ID: "123", Params: []Param{{Name: "foo:bar", Value: "high"}}}
		Assert(t, "did:example:123;foo:bar=high", d.String())
	})

	t.Run("includes Param generic and method", func(t *testing.T) {
		d := &DID{Method: "example", ID: "123",
			Params: []Param{{Name: "service", Value: "agent"}, {Name: "foo:bar", Value: "high"}}}
		Assert(t, "did:example:123;service=agent;foo:bar=high", d.String())
	})

	t.Run("includes Path", func(t *testing.T) {
		d := &DID{Method: "example", ID: "123", Path: "a/b"}
		Assert(t, "did:example:123/a/b", d.String())
	})

	t.Run("includes Path assembled from PathSegements", func(t *testing.T) {
		d := &DID{Method: "example", ID: "123", PathSegments: []string{"a", "b"}}
		Assert(t, "did:example:123/a/b", d.String())
	})

	t.Run("includes Path after Param", func(t *testing.T) {
		d := &DID{Method: "example", ID: "123",
			Params: []Param{{Name: "service", Value: "agent"}}, Path: "a/b"}
		Assert(t, "did:example:123;service=agent/a/b", d.String())
	})

	t.Run("includes Query after IDString", func(t *testing.T) {
		d := &DID{Method: "example", ID: "123", Query: "abc"}
		Assert(t, "did:example:123?abc", d.String())
	})

	t.Run("include Query after Param", func(t *testing.T) {
		d := &DID{Method: "example", ID: "123", Query: "abc",
			Params: []Param{{Name: "service", Value: "agent"}}}
		Assert(t, "did:example:123;service=agent?abc", d.String())
	})

	t.Run("includes Query after Path", func(t *testing.T) {
		d := &DID{Method: "example", ID: "123", Path: "x/y", Query: "abc"}
		Assert(t, "did:example:123/x/y?abc", d.String())
	})

	t.Run("includes Query after Param and Path", func(t *testing.T) {
		d := &DID{Method: "example", ID: "123", Path: "x/y", Query: "abc",
			Params: []Param{{Name: "service", Value: "agent"}}}
		Assert(t, "did:example:123;service=agent/x/y?abc", d.String())
	})

	t.Run("includes Query after before Fragment", func(t *testing.T) {
		d := &DID{Method: "example", ID: "123", Fragment: "zyx", Query: "abc"}
		Assert(t, "did:example:123?abc#zyx", d.String())
	})

	t.Run("includes Query", func(t *testing.T) {
		d := &DID{Method: "example", ID: "123", Query: "abc"}
		Assert(t, "did:example:123?abc", d.String())
	})

	t.Run("includes Fragment", func(t *testing.T) {
		d := &DID{Method: "example", ID: "123", Fragment: "00000"}
		Assert(t, "did:example:123#00000", d.String())
	})

	t.Run("includes Fragment after Param", func(t *testing.T) {
		d := &DID{Method: "example", ID: "123", Fragment: "00000"}
		Assert(t, "did:example:123#00000", d.String())
	})
}

func TestParse(t *testing.T) {

	t.Run("returns error if input is empty", func(t *testing.T) {
		_, err := Parse("")
		Assert(t, false, err == nil)
	})

	t.Run("returns error if input length is less than length 7", func(t *testing.T) {
		_, err := Parse("did:")
		Assert(t, false, err == nil)

		_, err = Parse("did:a")
		Assert(t, false, err == nil)

		_, err = Parse("did:a:")
		Assert(t, false, err == nil)
	})

	t.Run("returns error if input does not have a second : to mark end of method", func(t *testing.T) {
		_, err := Parse("did:aaaaaaaaaaa")
		Assert(t, false, err == nil)
	})

	t.Run("returns error if method is empty", func(t *testing.T) {
		_, err := Parse("did::aaaaaaaaaaa")
		Assert(t, false, err == nil)
	})

	t.Run("returns error if idstring is empty", func(t *testing.T) {
		dids := []string{
			"did:a::123:456",
			"did:a:123::456",
			"did:a:123:456:",
			"did:a:123:/abc",
			"did:a:123:#abc",
		}
		for _, did := range dids {
			_, err := Parse(did)
			Assert(t, false, err == nil, "Input: %s", did)
		}
	})

	t.Run("returns error if input does not begin with did: scheme", func(t *testing.T) {
		_, err := Parse("a:12345")
		Assert(t, false, err == nil)
	})

	t.Run("returned value is nil if input does not begin with did: scheme", func(t *testing.T) {
		d, _ := Parse("a:12345")
		Assert(t, true, d == nil)
	})

	t.Run("succeeds if it has did prefix and length is greater than 7", func(t *testing.T) {
		d, err := Parse("did:a:1")
		Assert(t, nil, err)
		Assert(t, true, d != nil)
	})

	t.Run("succeeds to extract method", func(t *testing.T) {
		d, err := Parse("did:a:1")
		Assert(t, nil, err)
		Assert(t, "a", d.Method)

		d, err = Parse("did:abcdef:11111")
		Assert(t, nil, err)
		Assert(t, "abcdef", d.Method)
	})

	t.Run("returns error if method has any other char than 0-9 or a-z", func(t *testing.T) {
		_, err := Parse("did:aA:1")
		Assert(t, false, err == nil)

		_, err = Parse("did:aa-aa:1")
		Assert(t, false, err == nil)
	})

	t.Run("succeeds to extract id", func(t *testing.T) {
		d, err := Parse("did:a:1")
		Assert(t, nil, err)
		Assert(t, "1", d.ID)
	})

	t.Run("succeeds to extract id parts", func(t *testing.T) {
		d, err := Parse("did:a:123:456")
		Assert(t, nil, err)

		parts := d.IDStrings
		Assert(t, "123", parts[0])
		Assert(t, "456", parts[1])
	})

	t.Run("returns error if ID has an invalid char", func(t *testing.T) {
		_, err := Parse("did:a:1&&111")
		Assert(t, false, err == nil)
	})

	t.Run("returns error if param name is empty", func(t *testing.T) {
		_, err := Parse("did:a:123:456;")
		Assert(t, false, err == nil)
	})

	t.Run("returns error if Param name has an invalid char", func(t *testing.T) {
		_, err := Parse("did:a:123:456;serv&ce")
		Assert(t, false, err == nil)
	})

	t.Run("returns error if Param value has an invalid char", func(t *testing.T) {
		_, err := Parse("did:a:123:456;service=ag&nt")
		Assert(t, false, err == nil)
	})

	t.Run("returns error if Param name has an invalid percent encoded", func(t *testing.T) {
		_, err := Parse("did:a:123:456;ser%2ge")
		Assert(t, false, err == nil)
	})

	t.Run("returns error if Param does not exist for value", func(t *testing.T) {
		_, err := Parse("did:a:123:456;=value")
		Assert(t, false, err == nil)
	})

	// nolint: dupl
	// test for params look similar to linter
	t.Run("succeeds to extract generic param with name and value", func(t *testing.T) {
		d, err := Parse("did:a:123:456;service==agent")
		Assert(t, nil, err)
		Assert(t, 1, len(d.Params))
		Assert(t, "service=agent", d.Params[0].String())
		Assert(t, "service", d.Params[0].Name)
		Assert(t, "agent", d.Params[0].Value)
	})

	// nolint: dupl
	// test for params look similar to linter
	t.Run("succeeds to extract generic param with name only", func(t *testing.T) {
		d, err := Parse("did:a:123:456;service")
		Assert(t, nil, err)
		Assert(t, 1, len(d.Params))
		Assert(t, "service", d.Params[0].String())
		Assert(t, "service", d.Params[0].Name)
		Assert(t, "", d.Params[0].Value)
	})

	// nolint: dupl
	// test for params look similar to linter
	t.Run("succeeds to extract generic param with name only and empty param", func(t *testing.T) {
		d, err := Parse("did:a:123:456;service=")
		Assert(t, nil, err)
		Assert(t, 1, len(d.Params))
		Assert(t, "service", d.Params[0].String())
		Assert(t, "service", d.Params[0].Name)
		Assert(t, "", d.Params[0].Value)
	})

	// nolint: dupl
	// test for params look similar to linter
	t.Run("succeeds to extract method param with name and value", func(t *testing.T) {
		d, err := Parse("did:a:123:456;foo:bar=baz")
		Assert(t, nil, err)
		Assert(t, 1, len(d.Params))
		Assert(t, "foo:bar=baz", d.Params[0].String())
		Assert(t, "foo:bar", d.Params[0].Name)
		Assert(t, "baz", d.Params[0].Value)
	})

	// nolint: dupl
	// test for params look similar to linter
	t.Run("succeeds to extract method param with name only", func(t *testing.T) {
		d, err := Parse("did:a:123:456;foo:bar")
		Assert(t, nil, err)
		Assert(t, 1, len(d.Params))
		Assert(t, "foo:bar", d.Params[0].String())
		Assert(t, "foo:bar", d.Params[0].Name)
		Assert(t, "", d.Params[0].Value)
	})

	// nolint: dupl
	// test for params look similar to linter
	t.Run("succeeds with percent encoded chars in param name and value", func(t *testing.T) {
		d, err := Parse("did:a:123:456;serv%20ice=val%20ue")
		Assert(t, nil, err)
		Assert(t, 1, len(d.Params))
		Assert(t, "serv%20ice=val%20ue", d.Params[0].String())
		Assert(t, "serv%20ice", d.Params[0].Name)
		Assert(t, "val%20ue", d.Params[0].Value)
	})

	// nolint: dupl
	// test for params look similar to linter
	t.Run("succeeds to extract multiple generic params with name only", func(t *testing.T) {
		d, err := Parse("did:a:123:456;foo;bar")
		Assert(t, nil, err)
		Assert(t, 2, len(d.Params))
		Assert(t, "foo", d.Params[0].Name)
		Assert(t, "", d.Params[0].Value)
		Assert(t, "bar", d.Params[1].Name)
		Assert(t, "", d.Params[1].Value)
	})

	// nolint: dupl
	// test for params look similar to linter
	t.Run("succeeds to extract multiple params with names and values", func(t *testing.T) {
		d, err := Parse("did:a:123:456;service=agent;foo:bar=baz")
		Assert(t, nil, err)
		Assert(t, 2, len(d.Params))
		Assert(t, "service", d.Params[0].Name)
		Assert(t, "agent", d.Params[0].Value)
		Assert(t, "foo:bar", d.Params[1].Name)
		Assert(t, "baz", d.Params[1].Value)
	})

	// nolint: dupl
	// test for params look similar to linter
	t.Run("succeeds to extract path after generic param", func(t *testing.T) {
		d, err := Parse("did:a:123:456;service==value/a/b")
		Assert(t, nil, err)
		Assert(t, 1, len(d.Params))
		Assert(t, "service=value", d.Params[0].String())
		Assert(t, "service", d.Params[0].Name)
		Assert(t, "value", d.Params[0].Value)

		segments := d.PathSegments
		Assert(t, "a", segments[0])
		Assert(t, "b", segments[1])
	})

	// nolint: dupl
	// test for params look similar to linter
	t.Run("succeeds to extract path after generic param name and no value", func(t *testing.T) {
		d, err := Parse("did:a:123:456;service=/a/b")
		Assert(t, nil, err)
		Assert(t, 1, len(d.Params))
		Assert(t, "service", d.Params[0].String())
		Assert(t, "service", d.Params[0].Name)
		Assert(t, "", d.Params[0].Value)

		segments := d.PathSegments
		Assert(t, "a", segments[0])
		Assert(t, "b", segments[1])
	})

	// nolint: dupl
	// test for params look similar to linter
	t.Run("succeeds to extract query after generic param", func(t *testing.T) {
		d, err := Parse("did:a:123:456;service=value?abc")
		Assert(t, nil, err)
		Assert(t, 1, len(d.Params))
		Assert(t, "service=value", d.Params[0].String())
		Assert(t, "service", d.Params[0].Name)
		Assert(t, "value", d.Params[0].Value)
		Assert(t, "abc", d.Query)
	})

	// nolint: dupl
	// test for params look similar to linter
	t.Run("succeeds to extract fragment after generic param", func(t *testing.T) {
		d, err := Parse("did:a:123:456;service=value#xyz")
		Assert(t, nil, err)
		Assert(t, 1, len(d.Params))
		Assert(t, "service=value", d.Params[0].String())
		Assert(t, "service", d.Params[0].Name)
		Assert(t, "value", d.Params[0].Value)
		Assert(t, "xyz", d.Fragment)
	})

	t.Run("succeeds to extract path", func(t *testing.T) {
		d, err := Parse("did:a:123:456/someService")
		Assert(t, nil, err)
		Assert(t, "someService", d.Path)
	})

	t.Run("succeeds to extract path segements", func(t *testing.T) {
		d, err := Parse("did:a:123:456/a/b")
		Assert(t, nil, err)

		segments := d.PathSegments
		Assert(t, "a", segments[0])
		Assert(t, "b", segments[1])
	})

	t.Run("succeeds with percent encoded chars in path", func(t *testing.T) {
		d, err := Parse("did:a:123:456/a/%20a")
		Assert(t, nil, err)
		Assert(t, "a/%20a", d.Path)
	})

	t.Run("returns error if % in path is not followed by 2 hex chars", func(t *testing.T) {
		dids := []string{
			"did:a:123:456/%",
			"did:a:123:456/%a",
			"did:a:123:456/%!*",
			"did:a:123:456/%A!",
			"did:xyz:pqr#%A!",
			"did:a:123:456/%A%",
		}
		for _, did := range dids {
			_, err := Parse(did)
			Assert(t, false, err == nil, "Input: %s", did)
		}
	})

	t.Run("returns error if path is empty but there is a slash", func(t *testing.T) {
		_, err := Parse("did:a:123:456/")
		Assert(t, false, err == nil)
	})

	t.Run("returns error if first path segment is empty", func(t *testing.T) {
		_, err := Parse("did:a:123:456//abc")
		Assert(t, false, err == nil)
	})

	t.Run("does not fail if second path segment is empty", func(t *testing.T) {
		_, err := Parse("did:a:123:456/abc//pqr")
		Assert(t, nil, err)
	})

	t.Run("returns error  if path has invalid char", func(t *testing.T) {
		_, err := Parse("did:a:123:456/ssss^sss")
		Assert(t, false, err == nil)
	})

	t.Run("does not fail if path has atleast one segment and a trailing slash", func(t *testing.T) {
		_, err := Parse("did:a:123:456/a/b/")
		Assert(t, nil, err)
	})

	t.Run("succeeds to extract query after idstring", func(t *testing.T) {
		d, err := Parse("did:a:123?abc")
		Assert(t, nil, err)
		Assert(t, "a", d.Method)
		Assert(t, "123", d.ID)
		Assert(t, "abc", d.Query)
	})

	t.Run("succeeds to extract query after path", func(t *testing.T) {
		d, err := Parse("did:a:123/a/b/c?abc")
		Assert(t, nil, err)
		Assert(t, "a", d.Method)
		Assert(t, "123", d.ID)
		Assert(t, "a/b/c", d.Path)
		Assert(t, "abc", d.Query)
	})

	t.Run("succeeds to extract fragment after query", func(t *testing.T) {
		d, err := Parse("did:a:123?abc#xyz")
		Assert(t, nil, err)
		Assert(t, "abc", d.Query)
		Assert(t, "xyz", d.Fragment)
	})

	t.Run("succeeds with percent encoded chars in query", func(t *testing.T) {
		d, err := Parse("did:a:123?ab%20c")
		Assert(t, nil, err)
		Assert(t, "ab%20c", d.Query)
	})

	t.Run("returns error if % in query is not followed by 2 hex chars", func(t *testing.T) {
		dids := []string{
			"did:a:123:456?%",
			"did:a:123:456?%a",
			"did:a:123:456?%!*",
			"did:a:123:456?%A!",
			"did:xyz:pqr?%A!",
			"did:a:123:456?%A%",
		}
		for _, did := range dids {
			_, err := Parse(did)
			Assert(t, false, err == nil, "Input: %s", did)
		}
	})

	t.Run("returns error if query has invalid char", func(t *testing.T) {
		_, err := Parse("did:a:123:456?ssss^sss")
		Assert(t, false, err == nil)
	})

	t.Run("succeeds to extract fragment", func(t *testing.T) {
		d, err := Parse("did:a:123:456#keys-1")
		Assert(t, nil, err)
		Assert(t, "keys-1", d.Fragment)
	})

	t.Run("succeeds with percent encoded chars in fragment", func(t *testing.T) {
		d, err := Parse("did:a:123:456#aaaaaa%20a")
		Assert(t, nil, err)
		Assert(t, "aaaaaa%20a", d.Fragment)
	})

	t.Run("returns error if % in fragment is not followed by 2 hex chars", func(t *testing.T) {
		dids := []string{
			"did:xyz:pqr#%",
			"did:xyz:pqr#%a",
			"did:xyz:pqr#%!*",
			"did:xyz:pqr#%!A",
			"did:xyz:pqr#%A!",
			"did:xyz:pqr#%A%",
		}
		for _, did := range dids {
			_, err := Parse(did)
			Assert(t, false, err == nil, "Input: %s", did)
		}
	})

	t.Run("fails if fragment has invalid char", func(t *testing.T) {
		_, err := Parse("did:a:123:456#ssss^sss")
		Assert(t, false, err == nil)
	})
}

func Test_errorf(t *testing.T) {
	p := &parser{}
	p.errorf(10, "%s,%s", "a", "b")

	if p.currentIndex != 10 {
		t.Errorf("did not set currentIndex")
	}

	e := p.err.Error()
	if e != "a,b" {
		t.Errorf("err message is: '%s' expected: 'a,b'", e)
	}
}

func Test_isNotValidParamChar(t *testing.T) {
	a := []byte{'A', 'B', 'C', 'D', 'E', 'F', 'G', 'H', 'I', 'J', 'K', 'L', 'M',
		'N', 'O', 'P', 'Q', 'R', 'S', 'T', 'U', 'V', 'W', 'X', 'Y', 'Z',
		'a', 'b', 'c', 'd', 'e', 'f', 'g', 'h', 'i', 'j', 'k', 'l', 'm',
		'n', 'o', 'p', 'q', 'r', 's', 't', 'u', 'v', 'w', 'x', 'y', 'z',
		'0', '1', '2', '3', '4', '5', '6', '7', '8', '9',
		'.', '-', '_', ':'}
	for _, c := range a {
		Assert(t, false, isNotValidParamChar(c), "Input: '%c'", c)
	}

	a = []byte{'%', '^', '#', ' ', '~', '!', '$', '&', '\'', '(', ')', '*', '+', ',', ';', '=', '@', '/', '?'}
	for _, c := range a {
		Assert(t, true, isNotValidParamChar(c), "Input: '%c'", c)
	}
}

func Test_isNotValidIDChar(t *testing.T) {
	a := []byte{'A', 'B', 'C', 'D', 'E', 'F', 'G', 'H', 'I', 'J', 'K', 'L', 'M',
		'N', 'O', 'P', 'Q', 'R', 'S', 'T', 'U', 'V', 'W', 'X', 'Y', 'Z',
		'a', 'b', 'c', 'd', 'e', 'f', 'g', 'h', 'i', 'j', 'k', 'l', 'm',
		'n', 'o', 'p', 'q', 'r', 's', 't', 'u', 'v', 'w', 'x', 'y', 'z',
		'0', '1', '2', '3', '4', '5', '6', '7', '8', '9',
		'.', '-'}
	for _, c := range a {
		Assert(t, false, isNotValidIDChar(c), "Input: '%c'", c)
	}

	a = []byte{'%', '^', '#', ' ', '_', '~', '!', '$', '&', '\'', '(', ')', '*', '+', ',', ';', '=', ':', '@', '/', '?'}
	for _, c := range a {
		Assert(t, true, isNotValidIDChar(c), "Input: '%c'", c)
	}
}

func Test_isNotValidQueryOrFragmentChar(t *testing.T) {
	a := []byte{'A', 'B', 'C', 'D', 'E', 'F', 'G', 'H', 'I', 'J', 'K', 'L', 'M',
		'N', 'O', 'P', 'Q', 'R', 'S', 'T', 'U', 'V', 'W', 'X', 'Y', 'Z',
		'a', 'b', 'c', 'd', 'e', 'f', 'g', 'h', 'i', 'j', 'k', 'l', 'm',
		'n', 'o', 'p', 'q', 'r', 's', 't', 'u', 'v', 'w', 'x', 'y', 'z',
		'0', '1', '2', '3', '4', '5', '6', '7', '8', '9',
		'-', '.', '_', '~', '!', '$', '&', '\'', '(', ')', '*', '+', ',', ';', '=',
		':', '@',
		'/', '?'}
	for _, c := range a {
		Assert(t, false, isNotValidQueryOrFragmentChar(c), "Input: '%c'", c)
	}

	a = []byte{'%', '^', '#', ' '}
	for _, c := range a {
		Assert(t, true, isNotValidQueryOrFragmentChar(c), "Input: '%c'", c)
	}
}

func Test_isNotValidPathChar(t *testing.T) {
	a := []byte{'A', 'B', 'C', 'D', 'E', 'F', 'G', 'H', 'I', 'J', 'K', 'L', 'M',
		'N', 'O', 'P', 'Q', 'R', 'S', 'T', 'U', 'V', 'W', 'X', 'Y', 'Z',
		'a', 'b', 'c', 'd', 'e', 'f', 'g', 'h', 'i', 'j', 'k', 'l', 'm',
		'n', 'o', 'p', 'q', 'r', 's', 't', 'u', 'v', 'w', 'x', 'y', 'z',
		'0', '1', '2', '3', '4', '5', '6', '7', '8', '9',
		'-', '.', '_', '~', '!', '$', '&', '\'', '(', ')', '*', '+', ',', ';', '=',
		':', '@'}
	for _, c := range a {
		Assert(t, false, isNotValidPathChar(c), "Input: '%c'", c)
	}

	a = []byte{'%', '/', '?'}
	for _, c := range a {
		Assert(t, true, isNotValidPathChar(c), "Input: '%c'", c)
	}
}

func Test_isNotUnreservedOrSubdelim(t *testing.T) {
	a := []byte{'A', 'B', 'C', 'D', 'E', 'F', 'G', 'H', 'I', 'J', 'K', 'L', 'M',
		'N', 'O', 'P', 'Q', 'R', 'S', 'T', 'U', 'V', 'W', 'X', 'Y', 'Z',
		'a', 'b', 'c', 'd', 'e', 'f', 'g', 'h', 'i', 'j', 'k', 'l', 'm',
		'n', 'o', 'p', 'q', 'r', 's', 't', 'u', 'v', 'w', 'x', 'y', 'z',
		'0', '1', '2', '3', '4', '5', '6', '7', '8', '9',
		'-', '.', '_', '~', '!', '$', '&', '\'', '(', ')', '*', '+', ',', ';', '='}
	for _, c := range a {
		Assert(t, false, isNotUnreservedOrSubdelim(c), "Input: '%c'", c)
	}

	a = []byte{'%', ':', '@', '/', '?'}
	for _, c := range a {
		Assert(t, true, isNotUnreservedOrSubdelim(c), "Input: '%c'", c)
	}
}

func Test_isNotHexDigit(t *testing.T) {
	a := []byte{'0', '1', '2', '3', '4', '5', '6', '7', '8', '9',
		'A', 'B', 'C', 'D', 'E', 'F', 'a', 'b', 'c', 'd', 'e', 'f'}
	for _, c := range a {
		Assert(t, false, isNotHexDigit(c), "Input: '%c'", c)
	}

	a = []byte{'G', 'g', '%', '\x40', '\x47', '\x60', '\x67'}
	for _, c := range a {
		Assert(t, true, isNotHexDigit(c), "Input: '%c'", c)
	}
}

func Test_isNotDigit(t *testing.T) {
	a := []byte{'0', '1', '2', '3', '4', '5', '6', '7', '8', '9'}
	for _, c := range a {
		Assert(t, false, isNotDigit(c), "Input: '%c'", c)
	}

	a = []byte{'A', 'a', '\x29', '\x40', '/'}
	for _, c := range a {
		Assert(t, true, isNotDigit(c), "Input: '%c'", c)
	}
}

func Test_isNotAlpha(t *testing.T) {
	a := []byte{'A', 'B', 'C', 'D', 'E', 'F', 'G', 'H', 'I', 'J', 'K', 'L', 'M',
		'N', 'O', 'P', 'Q', 'R', 'S', 'T', 'U', 'V', 'W', 'X', 'Y', 'Z',
		'a', 'b', 'c', 'd', 'e', 'f', 'g', 'h', 'i', 'j', 'k', 'l', 'm',
		'n', 'o', 'p', 'q', 'r', 's', 't', 'u', 'v', 'w', 'x', 'y', 'z'}
	for _, c := range a {
		Assert(t, false, isNotAlpha(c), "Input: '%c'", c)
	}

	a = []byte{'\x40', '\x5B', '\x60', '\x7B', '0', '9', '-', '%'}
	for _, c := range a {
		Assert(t, true, isNotAlpha(c), "Input: '%c'", c)
	}
}

// nolint: dupl
// Test_isNotSmallLetter and Test_isNotBigLetter look too similar to the dupl linter, ignore it
func Test_isNotBigLetter(t *testing.T) {
	a := []byte{'A', 'B', 'C', 'D', 'E', 'F', 'G', 'H', 'I', 'J', 'K', 'L', 'M',
		'N', 'O', 'P', 'Q', 'R', 'S', 'T', 'U', 'V', 'W', 'X', 'Y', 'Z'}
	for _, c := range a {
		Assert(t, false, isNotBigLetter(c), "Input: '%c'", c)
	}

	a = []byte{'\x40', '\x5B', 'a', 'z', '1', '9', '-', '%'}
	for _, c := range a {
		Assert(t, true, isNotBigLetter(c), "Input: '%c'", c)
	}
}

// nolint: dupl
// Test_isNotSmallLetter and Test_isNotBigLetter look too similar to the dupl linter, ignore it
func Test_isNotSmallLetter(t *testing.T) {
	a := []byte{'a', 'b', 'c', 'd', 'e', 'f', 'g', 'h', 'i', 'j', 'k', 'l', 'm',
		'n', 'o', 'p', 'q', 'r', 's', 't', 'u', 'v', 'w', 'x', 'y', 'z'}
	for _, c := range a {
		Assert(t, false, isNotSmallLetter(c), "Input: '%c'", c)
	}

	a = []byte{'\x60', '\x7B', 'A', 'Z', '1', '9', '-', '%'}
	for _, c := range a {
		Assert(t, true, isNotSmallLetter(c), "Input: '%c'", c)
	}
}

func Assert(t *testing.T, expected interface{}, actual interface{}, args ...interface{}) {
	if !reflect.DeepEqual(expected, actual) {
		argsLength := len(args)
		var message string

		// if only one arg is present, treat it as the message
		if argsLength == 1 {
			message = args[0].(string)
		}

		// if more than one arg is present, treat it as format, args (like Printf)
		if argsLength > 1 {
			message = fmt.Sprintf(args[0].(string), args[1:]...)
		}

		// is message is not empty add some spacing
		if message != "" {
			message = "\t" + message + "\n\n"
		}

		_, file, line, _ := runtime.Caller(1)
		fmt.Printf("%s:%d:\n\tExpected: %#v\n\tActual: %#v\n%s", filepath.Base(file), line, expected, actual, message)
		t.FailNow()
	}
}
