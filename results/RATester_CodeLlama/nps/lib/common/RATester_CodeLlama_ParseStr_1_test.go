package common

import (
	"fmt"
	"testing"
)

func TestParseStr_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	str := "{{.ENV_TEST}}"
	want := "test"
	got, err := ParseStr(str)
	if err != nil {
		t.Errorf("ParseStr() error = %v", err)
		return
	}
	if got != want {
		t.Errorf("ParseStr() = %v, want %v", got, want)
	}
}
