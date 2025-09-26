package accesslog

import (
	"fmt"
	"net/http"
	"testing"

	"github.com/sirupsen/logrus"
)

func TestRedactHeaders_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	// given
	h := &Handler{}
	headers := http.Header{}
	fields := logrus.Fields{}
	prefix := "prefix"

	// when
	h.redactHeaders(headers, fields, prefix)

	// then
	// TODO: add assertion
}
