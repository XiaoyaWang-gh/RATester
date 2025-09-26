package try

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"
	"testing"
)

func TestBodyContains_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	values := []string{"foo", "bar"}
	condition := BodyContains(values...)
	res := &http.Response{
		Body: ioutil.NopCloser(bytes.NewBufferString("foobar")),
	}
	if err := condition(res); err != nil {
		t.Errorf("expected no error, got %v", err)
	}
}
