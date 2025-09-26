package legacy

import (
	"fmt"
	"reflect"
	"testing"
)

func TestGetDefaultTokenConf_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	got := getDefaultTokenConf()
	want := TokenConfig{
		Token: "",
	}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("getDefaultTokenConf() = %v, want %v", got, want)
	}
}
