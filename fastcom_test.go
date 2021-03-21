package fastspeed

import (
	"os"
	"testing"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"os"
	"reflect"
	"testing"
	"time"

	sthttp "github.com/kylegrantlucas/speedtest/http"
)

func TestFastCom_Download(t *testing.T) {
	f, err := os.Open("http/random750x750.jpg")
	if err != nil {
		t.Fatal(err)
	}
	defer f.Close()

	b, err := ioutil.ReadAll(f)
	if err != nil {
		t.Fatal(err)
	}

	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, b)
	}))
	defer ts.Close()

	t.Run("testfastcom_download", func(t *testing.T)){
		got, err := tt.
	}
}
