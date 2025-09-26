package create

import (
	"fmt"
	"io/ioutil"
	"testing"
)

func TestBodyReader_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	o := fromRemoteOptions{
		Body: []byte("hello"),
	}
	r := o.BodyReader()
	if r == nil {
		t.Fatal("BodyReader() should not return nil")
	}
	b, err := ioutil.ReadAll(r)
	if err != nil {
		t.Fatal(err)
	}
	if string(b) != "hello" {
		t.Fatalf("BodyReader() should return %q, but got %q", "hello", string(b))
	}
}
