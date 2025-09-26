package orm

import (
	"fmt"
	"testing"
)

func TestOne_2(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	o := querySet{}
	container := struct{}{}
	err := o.One(container)
	if err != nil {
		t.Errorf("TestOne failed, expected nil, got %v", err)
	}
}
