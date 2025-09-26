package gin

import (
	"fmt"
	"io/ioutil"
	"mime/multipart"
	"net/http"
	"net/textproto"
	"os"
	"testing"
)

func TestSaveUploadedFile_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	file, err := os.Open("test.txt")
	if err != nil {
		t.Fatal(err)
	}
	defer file.Close()

	stat, err := file.Stat()
	if err != nil {
		t.Fatal(err)
	}

	header := &multipart.FileHeader{
		Filename: stat.Name(),
		Size:     stat.Size(),
		Header:   make(textproto.MIMEHeader),
	}

	tmpfile, err := ioutil.TempFile("", "upload_test")
	if err != nil {
		t.Fatal(err)
	}
	defer os.Remove(tmpfile.Name())

	ctx := &Context{
		Request: &http.Request{
			Method: "POST",
			Header: make(http.Header),
		},
	}

	err = ctx.SaveUploadedFile(header, tmpfile.Name())
	if err != nil {
		t.Fatal(err)
	}

	_, err = os.Stat(tmpfile.Name())
	if err != nil {
		t.Fatal(err)
	}
}
