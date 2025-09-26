package gin

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"
	"testing"
)

func TestMultipartForm_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()
	c := &Context{
		Request: &http.Request{
			Method: http.MethodPost,
			Header: http.Header{
				"Content-Type": []string{"multipart/form-data; boundary=----WebKitFormBoundary7MA4YWxkTrZu0gW"},
			},
			Body: ioutil.NopCloser(bytes.NewBufferString("------WebKitFormBoundary7MA4YWxkTrZu0gW\r\nContent-Disposition: form-data; name=\"name\"\r\n\r\njohn\r\n------WebKitFormBoundary7MA4YWxkTrZu0gW\r\nContent-Disposition: form-data; name=\"age\"\r\n\r\n25\r\n------WebKitFormBoundary7MA4YWxkTrZu0gW--\r\n")),
		},
	}
	form, err := c.MultipartForm()
	if err != nil {
		t.Errorf("MultipartForm() error = %v", err)
		return
	}
	if form == nil {
		t.Errorf("MultipartForm() form = nil")
		return
	}
	if form.Value["name"][0] != "john" {
		t.Errorf("MultipartForm() form.Value[\"name\"][0] = %v, want john", form.Value["name"][0])
	}
	if form.Value["age"][0] != "25" {
		t.Errorf("MultipartForm() form.Value[\"age\"][0] = %v, want 25", form.Value["age"][0])
	}
}
