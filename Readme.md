# fastspeed

Small library to test download/upload speed with fast.com and speedtest.net

## Example

```go
import (
	"github.com/Heto86/fastspeed"
)

func main() {
	// we have two options:
	// 	- fastspeed.FastcomType - to check fast.com
	//  - fastspeed.FastcomType - to check speedtest.net
	fastspeed.GetMeasurments(fastspeed.FastcomType)
	fastspeed.GetMeasurments(fastspeed.SpeednetType)
}
```
## Test

```sh
go test fastspee_test.go -v
```

## Benchmark

```sh
go test -bench=. 
```

## Dependencies

1. For fast.com we are using '''github.com/npotts/go-fast'''
1. For speedtest.net we are using '''github.com/kylegrantlucas/speedtest'''

## Known issues:

1. It looks that the upload for fast.com is not implemented in the library, unfortunately I couldn't find an half decent library except this.
1. The print output could be improved.
1. We could try to add more test servers.
