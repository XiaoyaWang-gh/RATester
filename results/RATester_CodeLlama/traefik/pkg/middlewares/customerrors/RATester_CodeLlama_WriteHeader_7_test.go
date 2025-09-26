package customerrors

import (
	"fmt"
	"testing"
)

func TestWriteHeader_7(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	r := &codeModifier{
		code: 200,
	}
	r.WriteHeader(200)
	if r.headerSent {
		t.Error("header should not be sent")
	}
	r.WriteHeader(100)
	if !r.headerSent {
		t.Error("header should be sent")
	}
}
