package httplib

import (
	"fmt"
	"io"
	"strings"
	"testing"
)

func TesthandleFiles_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	b := &BeegoHTTPRequest{
		url: "http://example.com",
		params: map[string][]string{
			"key1": {"value1"},
			"key2": {"value2"},
		},
		files: map[string]string{
			"file1": "/path/to/file1",
			"file2": "/path/to/file2",
		},
	}

	b.handleFiles()

	// Assert that the request body is as expected
	body, err := io.ReadAll(b.req.Body)
	if err != nil {
		t.Fatalf("Failed to read request body: %v", err)
	}

	// Assert that the content type is as expected
	contentType := b.req.Header.Get("Content-Type")
	if contentType != "multipart/form-data; boundary=..." {
		t.Errorf("Unexpected content type: %s", contentType)
	}

	// Assert that the request body contains the expected form fields
	if !strings.Contains(string(body), "key1=value1") {
		t.Errorf("Request body does not contain expected form field: key1=value1")
	}
	if !strings.Contains(string(body), "key2=value2") {
		t.Errorf("Request body does not contain expected form field: key2=value2")
	}

	// Assert that the request body contains the expected file fields
	if !strings.Contains(string(body), "file1=@/path/to/file1") {
		t.Errorf("Request body does not contain expected file field: file1=@/path/to/file1")
	}
	if !strings.Contains(string(body), "file2=@/path/to/file2") {
		t.Errorf("Request body does not contain expected file field: file2=@/path/to/file2")
	}
}
