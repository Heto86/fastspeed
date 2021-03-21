package main

import (
	"github.com/Heto86/fastspeed"
)

func main() {
	fastspeed.GetMeasurments(fastspeed.FastcomType)
	fastspeed.GetMeasurments(fastspeed.SpeednetType)
}
