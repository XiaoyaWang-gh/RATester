package common

import (
	"fmt"
	"testing"
)

func TestGetBoolByStr_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	if GetBoolByStr("1") != true {
		t.Errorf("GetBoolByStr(\"1\") = %v, want true", GetBoolByStr("1"))
	}
	if GetBoolByStr("true") != true {
		t.Errorf("GetBoolByStr(\"true\") = %v, want true", GetBoolByStr("true"))
	}
	if GetBoolByStr("0") != false {
		t.Errorf("GetBoolByStr(\"0\") = %v, want false", GetBoolByStr("0"))
	}
	if GetBoolByStr("false") != false {
		t.Errorf("GetBoolByStr(\"false\") = %v, want false", GetBoolByStr("false"))
	}
}
