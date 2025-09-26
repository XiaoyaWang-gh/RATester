package gin

import (
	"fmt"
	"io"
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

	fileHeader := &multipart.FileHeader{
		Filename: "test.txt",
		Size:     1024,
		Header:   make(textproto.MIMEHeader),
	}

	ctx := &Context{
		Request: &http.Request{
			Body: io.NopCloser(file),
		},
	}

	err = ctx.SaveUploadedFile(fileHeader, "test.txt")
	if err != nil {
		t.Fatal(err)
	}

	_, err = os.Stat("test.txt")
	if os.IsNotExist(err) {
		t.Fatal("File not created")
	}

	os.Remove("test.txt")
}
