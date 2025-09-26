package web

import (
	"fmt"
	"testing"
)

func TestBuildTemplate_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	var err error
	dir := "test"
	files := []string{"test.html"}
	err = BuildTemplate(dir, files...)
	if err != nil {
		t.Errorf("BuildTemplate() error = %v", err)
		return
	}
}
