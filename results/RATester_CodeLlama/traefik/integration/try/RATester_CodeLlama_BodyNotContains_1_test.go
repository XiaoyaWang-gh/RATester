package try

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"
	"testing"
)

func TestBodyNotContains_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	t.Parallel()

	values := []string{"foo", "bar"}
	condition := BodyNotContains(values...)

	res := &http.Response{
		Body: ioutil.NopCloser(bytes.NewBufferString("")),
	}

	err := condition(res)
	if err != nil {
		t.Errorf("expected no error, got %v", err)
	}
}
