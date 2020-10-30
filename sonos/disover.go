package sonos

import (
	"bufio"
	"net"
	"net/http"
	"strings"

	"github.com/LukasKnuth/sonos_ctrl/sonos/models"

	"github.com/reactivex/rxgo/v2"
)

const multicastAddr = "239.255.255.250:1900"
const discoverQuery = `M-SEARCH * HTTP/1.1
HOST: 239.255.255.250:1900
MAN: "ssdp:discover"
MX: 1
ST: urn:schemas-upnp-org:device:ZonePlayer:1`
const serverFilterWord = "Sonos"

// Discover Sonos speakers on the network automatically.
func Discover() <-chan rxgo.Item {
	out := make(chan rxgo.Item)

	go func() {
		conn, err := setupUDPDiscovery()
		if err != nil {
			out <- rxgo.Error(err)
		}
		reader := bufio.NewReader(conn)

		for true {
			req, err := http.ReadRequest(reader)
			if err != nil {
				out <- rxgo.Error(err)
			}
			ctrl, err := parseFoundHeader(req)
			if err != nil {
				out <- rxgo.Error(err)
			} else if ctrl != nil {
				out <- rxgo.Of(ctrl)
			}
		}
	}()

	return out
}

func parseFoundHeader(req *http.Request) (*models.Controller, error) {
	if server, ok := req.Header["Server"]; ok {
		if strings.Contains(server[0], serverFilterWord) {
			location := req.Header["Location"][0]
			usn := req.Header["Usn"][0]
			return models.ControllerFromDiscovery(location, usn)
		}
	}
	// Not a Sonos device
	return nil, nil
}

func setupUDPDiscovery() (*net.UDPConn, error) {
	addr, err := net.ResolveUDPAddr("udp4", multicastAddr)
	if err != nil {
		return nil, err
	}
	conn, err := net.ListenMulticastUDP("udp4", nil, addr)
	if err != nil {
		return nil, err
	}
	_, err = bufio.NewWriter(conn).WriteString(discoverQuery)
	if err != nil {
		return nil, err
	}
	return conn, nil
}
