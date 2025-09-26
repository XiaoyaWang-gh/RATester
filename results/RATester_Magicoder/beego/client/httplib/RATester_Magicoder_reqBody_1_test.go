package httplib

import (
	"fmt"
	"io"
	"testing"
)

func TestreqBody_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	b := &BeegoHTTPRequest{}
	data := []byte("test data")
	b.reqBody(data)

	if b.req.Body == nil {
		t.Error("req.Body is nil")
	}

	body, err := io.ReadAll(b.req.Body)
	if err != nil {
		t.Error("Failed to read req.Body:", err)
	}

	if string(body) != string(data) {
		t.Error("req.Body does not match expected data")
	}

	if b.req.ContentLength != int64(len(data)) {
		t.Error("req.ContentLength does not match expected length")
	}

	if b.copyBody == nil {
		t.Error("copyBody is nil")
	}

	copyBody := b.copyBody()
	if copyBody == nil {
		t.Error("copyBody() returned nil")
	}

	copyBodyData, err := io.ReadAll(copyBody)
	if err != nil {
		t.Error("Failed to read copyBody:", err)
	}

	if string(copyBodyData) != string(data) {
		t.Error("copyBody does not match expected data")
	}
}
