package fastspeed

//this library has a good quality
import (
	gofast "github.com/npotts/go-fast"
)

// we could add things here like gf or cfg, but makes little sense
type FastCom struct {
}

// to keep the consistency with speedtestnet
func (f FastCom) Init() FastCom {
	return f
}

// theoretically the bool should change the direction, it is not the case
func (f FastCom) measure(emit_json bool) float64 {
	gf := gofast.New()
	cfg := gofast.Settings{
		MaxBytes: int64(0),
		Timeout:  0,
		Workers:  int(3),
		Network:  5000000000,
		EmitJSON: emit_json}
	results := <-gf.Measure(cfg)
	return results.Mbps
}

func (f FastCom) Download() float64 {
	return f.measure(false)
}

// i'm not really sure that this is really doing the upload, but I would rather modify
// the library than this part of code
func (f FastCom) Upload() float64 {
	return f.measure(true)
}
