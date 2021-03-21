package fastspeed

import (
	"fmt"

	"github.com/kylegrantlucas/speedtest"
)

type SpeedTestNet struct {
	client *speedtest.Client
}

func (s SpeedTestNet) Init() SpeedTestNet {
	client, err := speedtest.NewDefaultClient()
	if err != nil {
		fmt.Printf("error creating client: %v", err)
	}
	s.client = client
	return s
}

func (s SpeedTestNet) Download() float64 {
	server, err := s.client.GetServer("")
	if err != nil {
		fmt.Printf("error getting server: %v", err)
	}
	dmbps, err := s.client.Download(server)
	if err != nil {
		fmt.Printf("error getting download: %v", err)
	}
	return dmbps
}

func (s SpeedTestNet) Upload() float64 {
	server, err := s.client.GetServer("")
	if err != nil {
		fmt.Printf("error getting download: %v", err)
	}
	dmbps, err := s.client.Upload(server)
	if err != nil {
		fmt.Printf("error getting server: %v", err)
	}
	return dmbps
}
