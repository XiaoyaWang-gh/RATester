package try

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
	"testing"
)

func TestBodyContainsOr_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	values := []string{"foo", "bar"}
	condition := BodyContainsOr(values...)
	res := &http.Response{
		Body: ioutil.NopCloser(strings.NewReader("foobar")),
	}
	if err := condition(res); err != nil {
		t.Errorf("expected no error, got %v", err)
	}
}
