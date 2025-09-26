package client

import (
	"fmt"
	"testing"
)

func TestFile_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	r := &Request{}
	r.AddFile("test.txt")
	r.AddFileWithReader("test.txt", nil)
	r.AddFiles(nil)
	r.File("test.txt")
	r.FileByPath("test.txt")
	r.FormData("test.txt")
	r.Get("test.txt")
	r.Head("test.txt")
	r.Header("test.txt")
	r.Param("test.txt")
	r.Post("test.txt")
	r.Put("test.txt")
	r.Patch("test.txt")
	r.Options("test.txt")
	r.Delete("test.txt")
	r.SetFormData("test.txt", "test.txt")
	r.SetFormDatas(map[string]string{"test.txt": "test.txt"})
	r.SetFormDatasWithStruct(struct{}{})
}
