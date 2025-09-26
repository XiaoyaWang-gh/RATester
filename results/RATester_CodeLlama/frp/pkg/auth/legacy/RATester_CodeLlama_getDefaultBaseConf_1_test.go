package legacy

import (
	"fmt"
	"reflect"
	"testing"
)

func TestGetDefaultBaseConf_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	want := BaseConfig{
		AuthenticationMethod:     "token",
		AuthenticateHeartBeats:   false,
		AuthenticateNewWorkConns: false,
	}
	got := getDefaultBaseConf()
	if !reflect.DeepEqual(want, got) {
		t.Errorf("getDefaultBaseConf() = %v, want %v", got, want)
	}
}
