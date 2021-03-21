package fastspeed

import (
	"bufio"
	"os"
	"regexp"
	"testing"
	//"github.com/Heto86/fastspeed"
)

func TestSpeedTestNet_Download(t *testing.T) {
	t.Logf("testing speedtest.net download")
	var f SpeedTestNet
	f = f.Init()
	dlSpeed := Download(f)
	if 15.0 > dlSpeed || dlSpeed > 120.0 {
		t.Fail()
	}
}

func TestSpeedTestNet_Upload(t *testing.T) {
	t.Logf("testing speedtest.net upload")
	var f SpeedTestNet
	f = f.Init()
	ulSpeed := Upload(f)
	if 15.0 > ulSpeed || ulSpeed > 120.0 {
		t.Fail()
	}
}

func TestFastCom_Download(t *testing.T) {
	t.Logf("testing fast.com download")
	var f FastCom
	f = f.Init()
	dlSpeed := Download(f)
	if 15.0 > dlSpeed || dlSpeed > 120.0 {
		t.Fail()
	}
}

func TestFastCom_Upload(t *testing.T) {
	t.Logf("testing fast.com upload")
	var f FastCom
	f = f.Init()
	ulSpeed := Upload(f)
	if 15.0 > ulSpeed || ulSpeed > 120.0 {
		t.Fail()
	}
}

func TestAPIFastCom_GetMeasurments(t *testing.T) {
	t.Logf("testing API GetMeasurments with FastcomType")

	old_stdout := os.Stdout
	reader, writer, err := os.Pipe()
	if err != nil {
		t.Fail()
	}
	os.Stdout = writer
	scanner := bufio.NewScanner(reader)
	var outputs []string
	go func() {
		for scanner.Scan() {
			line := scanner.Text()
			outputs = append(outputs, line)
		}
	}()
	GetMeasurments(FastcomType)
	writer.Close()
	os.Stdout = old_stdout // restoring the real stdout
	for _, output := range outputs {
		re_output := regexp.MustCompile(`(Upload|Download): \d{2,3}\.\d{2}`)
		if !re_output.MatchString(output) {
			t.Fail()
		}
	}
}

func TestAPISpeedTestNet_GetMeasurments(t *testing.T) {
	t.Logf("testing API GetMeasurments with SpeedtestnetType")

	old_stdout := os.Stdout
	reader, writer, err := os.Pipe()
	if err != nil {
		t.Fail()
	}
	os.Stdout = writer
	scanner := bufio.NewScanner(reader)
	var outputs []string
	go func() {
		for scanner.Scan() {
			line := scanner.Text()
			outputs = append(outputs, line)
		}
	}()
	GetMeasurments(SpeednetType)
	writer.Close()
	os.Stdout = old_stdout // restoring the real stdout
	for _, output := range outputs {
		re_output := regexp.MustCompile(`(Upload|Download): \d{2,3}\.\d{2}`)
		if !re_output.MatchString(output) {
			t.Fail()
		}
	}
}

func Benchmark_DownloadFastcom(b *testing.B) {
	for n := 0; n < b.N; n++ {
		var f FastCom
		f = f.Init()
		Download(f)
	}
}

func Benchmark_UploadFastcom(b *testing.B) {
	for n := 0; n < b.N; n++ {
		var f FastCom
		f = f.Init()
		Upload(f)
	}
}

func Benchmark_DownloadSpeedtestnet(b *testing.B) {
	for n := 0; n < b.N; n++ {
		var f SpeedTestNet
		f = f.Init()
		Download(f)
	}
}

func Benchmark_UploadSpeedtestnet(b *testing.B) {
	for n := 0; n < b.N; n++ {
		var f SpeedTestNet
		f = f.Init()
		Download(f)
	}
}

func BenchmarkFastcom_GetMeasurments(b *testing.B) {
	for n := 0; n < b.N; n++ {
		GetMeasurments(FastcomType)
	}
}

func BenchmarkSpeedtestnet_GetMeasurments(b *testing.B) {
	for n := 0; n < b.N; n++ {
		GetMeasurments(SpeednetType)
	}
}
