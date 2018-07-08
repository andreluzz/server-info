package server_info

import (
	"os"
	"net"
	"strings"
	"errors"
	"runtime"
	"time"
	"encoding/json"
)

type Adapter struct {
	Name string	`json:"name"`
	IP string	`json:"ip"`
	MAC string	`json:"mac"`
}

type Server struct {
	Hostname string 	`json:"hostname"`
	OS string			`json:"os"`
	Adapters []Adapter	`json:"adapters"`
	Date time.Time		`json:"date"`
}

func New() *Server {
	server := &Server{}
	return server
}

func (server *Server) Flush() {
	name, _ := os.Hostname()
	server.Hostname = name

	server.OS = runtime.GOOS

	networkInterfaces, _ := net.Interfaces()
	server.Adapters = nil
	for _, i := range networkInterfaces {
		if strings.Contains(i.Flags.String(), "up") {
			addresses, _ := i.Addrs()
			for _, a := range  addresses {
				if interfaceIP, ok := a.(*net.IPNet); ok && !interfaceIP.IP.IsLoopback() {
					adapter := Adapter{
						i.Name,
						interfaceIP.IP.To4().String(),
						i.HardwareAddr.String()}
					server.Adapters = append(server.Adapters, adapter)
				}
			}
		}
	}
}

func (server *Server) GetIPAddress(adapterName ...string) (string, error) {
	if len(server.Adapters) == 0 {
		return "", errors.New("no network adapter up found")
	}

	if len(adapterName) == 0 {
		return server.Adapters[0].IP, nil
	}

	for _, a := range server.Adapters {
		if a.Name == adapterName[0] {
			return a.IP, nil
		}
	}

	return "", errors.New("no network adapter found")
}

func (server *Server) ToJSON() ([]byte, error) {
	server.Date = time.Now()
	return json.Marshal(server)
}
