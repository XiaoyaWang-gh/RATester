package config

import (
	"fmt"
	"testing"
)

func TestParseData_2(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	ini := &IniConfig{}
	data := []byte("key=value")
	_, err := ini.ParseData(data)
	if err != nil {
		t.Errorf("ParseData() error = %v, wantErr %v", err, false)
		return
	}
}
