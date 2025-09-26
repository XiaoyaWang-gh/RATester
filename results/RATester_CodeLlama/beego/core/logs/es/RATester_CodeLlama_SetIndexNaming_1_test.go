package es

import (
	"fmt"
	"testing"
)

func TestSetIndexNaming_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	var i IndexNaming
	SetIndexNaming(i)
	if indexNaming != i {
		t.Errorf("indexNaming is not i")
	}
}
