package web

import (
	"fmt"
	"mime/multipart"
	"net/http"
	"testing"

	"github.com/beego/beego/v2/server/web/context"
)

func TestGetFiles_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	ctrl := &Controller{
		Ctx: &context.Context{
			Request: &http.Request{
				MultipartForm: &multipart.Form{
					File: map[string][]*multipart.FileHeader{
						"key": {
							{
								Filename: "file1.txt",
							},
							{
								Filename: "file2.txt",
							},
						},
					},
				},
			},
		},
	}

	files, err := ctrl.GetFiles("key")
	if err != nil {
		t.Errorf("Expected no error, but got %v", err)
	}

	if len(files) != 2 {
		t.Errorf("Expected 2 files, but got %d", len(files))
	}

	if files[0].Filename != "file1.txt" {
		t.Errorf("Expected file1.txt, but got %s", files[0].Filename)
	}

	if files[1].Filename != "file2.txt" {
		t.Errorf("Expected file2.txt, but got %s", files[1].Filename)
	}
}
