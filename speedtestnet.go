package fastspeed

import (
	"fmt"

	"github.com/kylegrantlucas/speedtest"
	"github.com/kylegrantlucas/speedtest/http"
)

type SpeedTestNet struct {
	client *speedtest.Client
	server http.Server
}

func (s SpeedTestNet) Init() SpeedTestNet {
	client, err := speedtest.NewDefaultClient()
	if err != nil {
		fmt.Printf("error creating client: %v", err)
	}
	s.client = client

	server, err := s.client.GetServer("")
	if err != nil {
		fmt.Printf("error getting server: %v", err)
	}
	s.server = server
	return s
}

func (s SpeedTestNet) Download() float64 {
	dmbps, err := s.client.Download(s.server)
	if err != nil {
		fmt.Printf("error download : %v", err)
	}
	return dmbps
}

func (s SpeedTestNet) Upload() float64 {
	dmbps, err := s.client.Upload(s.server)
	if err != nil {
		fmt.Printf("error upload: %v", err)
	}
	return dmbps
}
