package fastspeed

// excellent library
import (
	"fmt"

	"github.com/kylegrantlucas/speedtest"
	"github.com/kylegrantlucas/speedtest/http"
)

// to be as dry as we can, we keep the client and the server here
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

// Self explanatory
func (s SpeedTestNet) Download() float64 {
	dmbps, err := s.client.Download(s.server)
	if err != nil {
		fmt.Printf("error download : %v", err)
	}
	return dmbps
}

// Self explanatory
func (s SpeedTestNet) Upload() float64 {
	umbps, err := s.client.Upload(s.server)
	if err != nil {
		fmt.Printf("error upload: %v", err)
	}
	return umbps
}
