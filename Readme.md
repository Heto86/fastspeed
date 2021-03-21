# fastspeed
[![license](https://img.shields.io/github/license/DAVFoundation/captain-n3m0.svg?style=flat-square)](https://github.com/DAVFoundation/captain-n3m0/blob/master/LICENSE)

Small library to test download/upload speed with fast.com and speedtest.net

## Example

```go
import (
	"github.com/Heto86/fastspeed"
)

func main() {
	// we have two options:
	// 	- fastspeed.FastcomType - to check fast.com
	//  	- fastspeed.FastcomType - to check speedtest.net
	fastspeed.GetMeasurments(fastspeed.FastcomType)
	fastspeed.GetMeasurments(fastspeed.SpeednetType)
}
```
## Test

```sh
go test fastspeed_test.go -v
```

## Benchmark

```sh
go test -bench=. 
```

## Dependencies

1. For fast.com we are using '''github.com/npotts/go-fast'''
1. For speedtest.net we are using '''github.com/kylegrantlucas/speedtest'''

Full list of dependencies:
- github.com/dchest/uniuri
- github.com/kylegrantlucas/speedtest
- github.com/kylegrantlucas/speedtest/coords
- github.com/kylegrantlucas/speedtest/http
- github.com/kylegrantlucas/speedtest/util
- github.com/kylegrantlucas/speedtest/xml
- github.com/npotts/go-fast
- github.com/pkg/errors
- golang.org/x/net/html
- golang.org/x/net/html/atom

## Known issues:

1. It looks that the upload for fast.com is not implemented in the library.
1. The print output could be improved.
1. We could try to add more test servers.
