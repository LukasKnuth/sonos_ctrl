package models

import (
	"fmt"
	"net"
	"net/url"
)

// Identifies a single Sonos controller (either a single Speaker or the leader of a Group)
type Controller struct {
	// IPv4 address of this controller
	IP string
	// Uinique identifier for this specific hardware
	USN string
	// The location to discover this device, as returned by UPNP lookup
	Location string
}

// Creates a new Controller model from information aquired via UPNP discovery
func ControllerFromDiscovery(location string, usn string) (*Controller, error) {
	ip, err := ipFromLocation(location)
	if err != nil {
		return nil, err
	}
	return &Controller{IP: ip, USN: usn, Location: location}, nil
}

func (ctrl Controller) String() string {
	return fmt.Sprintf("SonosController{IP: %v, USN: %v}", ctrl.IP, ctrl.USN)
}

func ipFromLocation(location string) (string, error) {
	url, err := url.Parse(location)
	if err != nil {
		return "", err
	} else {
		host, _, err := net.SplitHostPort(url.Host)
		if err != nil {
			return "", err
		}
		return host, nil
	}
}
