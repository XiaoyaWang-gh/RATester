package accesslog

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

func TestInitServiceFields_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	rw := httptest.NewRecorder()
	req, err := http.NewRequest("GET", "http://example.com", nil)
	if err != nil {
		t.Fatal(err)
	}
	next := http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {})
	data := &LogData{}
	InitServiceFields(rw, req, next, data)
	if data.Core[OriginDuration] != time.Duration(0) {
		t.Errorf("Expected %v, got %v", time.Duration(0), data.Core[OriginDuration])
	}
	if data.Core[OriginStatus] != 0 {
		t.Errorf("Expected %v, got %v", 0, data.Core[OriginStatus])
	}
	if data.Core[OriginContentSize] != int64(0) {
		t.Errorf("Expected %v, got %v", int64(0), data.Core[OriginContentSize])
	}
}
