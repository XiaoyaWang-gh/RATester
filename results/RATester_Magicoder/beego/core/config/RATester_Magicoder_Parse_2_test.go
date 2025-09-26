package config

import (
	"fmt"
	"testing"
)

func TestParse_2(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	ini := &IniConfig{}
	_, err := ini.Parse("test.ini")
	if err != nil {
		t.Errorf("Parse failed, err: %v", err)
	}
}
