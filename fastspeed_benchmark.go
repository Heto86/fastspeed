package fastspeed

import (
	"testing"
)

func BenchMark_DownloadFastcom(b *testing.B) {
	for n := 0; n < b.N; n++ {
		var f FastCom
		f = f.Init()
		Download(f)
	}
}

func BenchMark_UploadFastcom(b *testing.B) {
	for n := 0; n < b.N; n++ {
		var f FastCom
		f = f.Init()
		Upload(f)
	}
}

func BenchMark_DownloadSpeedtestnet(b *testing.B) {
	for n := 0; n < b.N; n++ {
		var f SpeedTestNet
		f = f.Init()
		Download(f)
	}
}

func BenchMark_UploadSpeedtestnet(b *testing.B) {
	for n := 0; n < b.N; n++ {
		var f SpeedTestNet
		f = f.Init()
		Download(f)
	}
}

func BenchMarkFastcom_GetMeasurments(b *testing.B) {
	for n := 0; n < b.N; n++ {
		GetMeasurments(FastcomType)
	}
}

func BenchMarkSpeedtestnet_GetMeasurments(b *testing.B) {
	for n := 0; n < b.N; n++ {
		GetMeasurments(SpeednetType)
	}
}
