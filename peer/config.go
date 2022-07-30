package peer

import (
	"fmt"
	"net"

	"github.com/hyperledger/fabric/common/flogging"
	"github.com/pkg/errors"
	"github.com/spf13/viper"
)

var peerLogger = flogging.MustGetLogger("peer")

type ServerConfig struct {
}

type ClientConfig struct {
}

type Config struct {
	LocalMSPID string
	ListenAddr string
	PeerID     string
	PeerAddr   string
	NetworkID  string
	TLSConfig  TLSConfig
}

type TLSConfig struct {
	Enabled     bool
	PrivKeyFile string
	CertFile    string
}

func (c *Config) load() error {
	peerAddr, err := getLocalAddress()
	c.PeerAddr = peerAddr
	return err
}

// getLocalAddress returns the address:port the local peer is operating on.  Affected by env:peer.addressAutoDetect
func getLocalAddress() (string, error) {
	peerAddress := viper.GetString("peer.address")
	if peerAddress == "" {
		return "", fmt.Errorf("peer.address isn't set")
	}
	host, port, err := net.SplitHostPort(peerAddress)
	if err != nil {
		return "", errors.Errorf("peer.address isn't in host:port format: %s", peerAddress)
	}

	localIP, err := getLocalIP()
	if err != nil {
		peerLogger.Errorf("local IP address not auto-detectable: %s", err)
		return "", err
	}
	autoDetectedIPAndPort := net.JoinHostPort(localIP, port)
	peerLogger.Info("Auto-detected peer address:", autoDetectedIPAndPort)
	// If host is the IPv4 address "0.0.0.0" or the IPv6 address "::",
	// then fallback to auto-detected address
	if ip := net.ParseIP(host); ip != nil && ip.IsUnspecified() {
		peerLogger.Info("Host is", host, ", falling back to auto-detected address:", autoDetectedIPAndPort)
		return autoDetectedIPAndPort, nil
	}

	if viper.GetBool("peer.addressAutoDetect") {
		peerLogger.Info("Auto-detect flag is set, returning", autoDetectedIPAndPort)
		return autoDetectedIPAndPort, nil
	}
	peerLogger.Info("Returning", peerAddress)
	return peerAddress, nil
}

// getLocalIP returns the a loopback local IP of the host.
func getLocalIP() (string, error) {
	addrs, err := net.InterfaceAddrs()
	if err != nil {
		return "", err
	}
	for _, address := range addrs {
		// check the address type and if it is not a loopback then display it
		if ipnet, ok := address.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
			if ipnet.IP.To4() != nil {
				return ipnet.IP.String(), nil
			}
		}
	}
	return "", errors.Errorf("no non-loopback, IPv4 interface detected")
}
