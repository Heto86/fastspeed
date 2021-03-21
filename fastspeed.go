package fastspeed

import "fmt"

// Both Fast.com and Speedtest.net implement this interface
type FastSpeed interface {
	Download() float64
	Upload() float64
}

func Download(fs FastSpeed) float64 {
	return fs.Download()
}

func Upload(fs FastSpeed) float64 {
	return fs.Upload()
}

// We can add any number of new "servers".
type FastspeedType int

const (
	FastcomType = iota
	SpeednetType
)

// This is the function that is given to outside world
func GetMeasurments(ft FastspeedType) {
	fast_speed := create_fastspeed(ft)
	dmbs, umbs := measure(fast_speed)
	print_measurments(dmbs, umbs)
}

// Only two servers, if - else if - else is good enough and is
// clear what is happening. If the number of servers will rise, we can
// take advantage of the fact that the enums could be threated as indexes
// and we can create a table with types, retrieving the value form the given index
func create_fastspeed(ft FastspeedType) FastSpeed {
	if ft == SpeednetType {
		var speednet SpeedTestNet
		return speednet.Init()
	} else if ft == FastcomType {
		var fastcom FastCom
		return fastcom.Init()
	} else {
		panic("SpeedTest provider is unknown!")
	}
}

// self explenatory
func measure(fast_speed FastSpeed) (float64, float64) {
	dmbs := Download(fast_speed)
	umbs := Upload(fast_speed)
	return dmbs, umbs
}

// self explenatory
func print_measurments(dmbs, umbs float64) {
	fmt.Printf("Download: %f2\n", dmbs)
	fmt.Printf("Upload: %f2\n", umbs)
}
