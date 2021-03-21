package fastspeed

import "fmt"

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

type FastspeedType int

const (
	FastcomType = iota
	SpeednetType
)

func GetMeasurments(ft FastspeedType) {
	fast_speed := create_fastspeed(ft)
	dmbs, umbs := measure(fast_speed)
	print_measurments(dmbs, umbs)
}

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

func measure(fast_speed FastSpeed) (float64, float64) {
	dmbs := Download(fast_speed)
	umbs := Upload(fast_speed)
	return dmbs, umbs
}

func print_measurments(dmbs, umbs float64) {
	fmt.Printf("Download: %f2\n", dmbs)
	fmt.Printf("Upload: %f2\n", umbs)
}
