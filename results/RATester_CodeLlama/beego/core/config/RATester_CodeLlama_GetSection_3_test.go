package config

import (
	"fmt"
	"testing"
)

func TestGetSection_3(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	var err error
	var section string
	var result map[string]string

	section = "section1"
	result, err = GetSection(section)
	if err != nil {
		t.Errorf("GetSection failed, err: %v", err)
		return
	}
	if len(result) != 2 {
		t.Errorf("GetSection failed, result: %v", result)
		return
	}

	section = "section2"
	result, err = GetSection(section)
	if err != nil {
		t.Errorf("GetSection failed, err: %v", err)
		return
	}
	if len(result) != 3 {
		t.Errorf("GetSection failed, result: %v", result)
		return
	}
}
