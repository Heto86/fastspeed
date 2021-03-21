package main

import (
	"fmt"

	"github.com/Heto86/fastspeed"
)

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

func main() {
	var f fastspeed.FastCom
	dlSpeed := Download(f)
	fmt.Printf("%f", dlSpeed)
}
