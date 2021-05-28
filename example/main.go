package main

import (
	"github.com/Heto86/fastspeed"
)

func main() {
	// we have two options:
	//  - fastspeed.FastcomType - to check fast.com
	//  - fastspeed.SpeednetType - to check speedtest.net
	fastspeed.GetMeasurments(fastspeed.FastcomType)
	fastspeed.GetMeasurments(fastspeed.SpeednetType)
}
