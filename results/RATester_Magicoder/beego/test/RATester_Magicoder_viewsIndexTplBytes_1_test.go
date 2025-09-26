package test

import (
	"fmt"
	"testing"
)

func TestviewsIndexTplBytes_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	data, err := viewsIndexTplBytes()
	if err != nil {
		t.Errorf("viewsIndexTplBytes() returned an error: %v", err)
	}

	if len(data) == 0 {
		t.Error("viewsIndexTplBytes() returned an empty byte slice")
	}
}
