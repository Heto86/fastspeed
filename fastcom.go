package fastspeed

import (
	gofast "github.com/npotts/go-fast"
)

type FastCom struct {
}

func (f FastCom) Init() FastCom {
	return f
}

func (f FastCom) measure(emit_json bool) float64 {
	gf := gofast.New()
	cfg := gofast.Settings{MaxBytes: int64(0), Timeout: 0, Workers: int(3), Network: 5000000000, EmitJSON: emit_json}
	results := <-gf.Measure(cfg)
	return results.Mbps
}

func (f FastCom) Download() float64 {
	return f.measure(false)
}

func (f FastCom) Upload() float64 {
	return f.measure(true)
}
