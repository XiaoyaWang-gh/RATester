package http

import (
	"fmt"
	"testing"
)

func TestOkResponse_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	res := OkResponse()
	if res.Status != "OK" {
		t.Errorf("Status = %q; want \"OK\"", res.Status)
	}
	if res.StatusCode != 200 {
		t.Errorf("StatusCode = %d; want 200", res.StatusCode)
	}
	if res.Proto != "HTTP/1.1" {
		t.Errorf("Proto = %q; want \"HTTP/1.1\"", res.Proto)
	}
	if res.ProtoMajor != 1 {
		t.Errorf("ProtoMajor = %d; want 1", res.ProtoMajor)
	}
	if res.ProtoMinor != 1 {
		t.Errorf("ProtoMinor = %d; want 1", res.ProtoMinor)
	}
}
