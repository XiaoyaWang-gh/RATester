package httplib

import (
	"fmt"
	"reflect"
	"testing"
)

func TestBytes_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	// TODO: setup test
	b := &BeegoHTTPRequest{}
	gotBytes, err := b.Bytes()
	if err != nil {
		t.Errorf("Bytes() error = %v", err)
		return
	}
	if !reflect.DeepEqual(gotBytes, []byte{}) {
		t.Errorf("Bytes() = %v, want %v", gotBytes, []byte{})
	}
}
