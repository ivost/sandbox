package server

import (
	"context"
	"net"
	"strings"
	"time"

	// client of myservice
	mscl "github.com/ivost/sandbox/myservice/client"
	mscfg "github.com/ivost/sandbox/myservice/config"

	"github.com/golang/protobuf/ptypes/empty"
	v1 "github.com/ivost/sandbox/myvault/myvault"
	"github.com/ivost/sandbox/myvault/pkg/version"
)

func (s *Server) Health(ctx context.Context, none *empty.Empty) (resp *v1.HealthResponse, err error) {
	resp = &v1.HealthResponse{
		Name:    "myvault",
		Version: version.Version,
		Build:   version.Build,
		Status:  "OK",
		Time:    time.Now().Format(time.RFC3339),
		Address: MyIP(),
	}

	// call myservice Health
	cfg := mscfg.DefaultConfig()
	cl := mscl.New(cfg)
	rsp :=
	resp.Status := "myclient: " + rsp.
	return
}

func MyIP() string {
	ifaces, err := net.Interfaces()
	// handle err
	_ = err
	for _, i := range ifaces {
		addrs, err := i.Addrs()
		// handle err
		_ = err
		for _, addr := range addrs {
			var ip net.IP
			switch v := addr.(type) {
			case *net.IPNet:
				ip = v.IP
			case *net.IPAddr:
				ip = v.IP
			}
			if ip == nil {
				continue
			}
			s := ip.String()
			if strings.Contains(s, ":") {
				continue
			}
			if s == "127.0.0.1" {
				continue
			}
			//log.Printf("addr: %v", ip)
			return s
		}
	}
	return ""
}
