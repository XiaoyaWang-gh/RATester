package test

import (
	"fmt"
	"testing"
)

func TestviewsHeaderTplBytes_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	_, err := viewsHeaderTplBytes()
	if err != nil {
		t.Errorf("viewsHeaderTplBytes() returned an error: %v", err)
	}
}
