package httplib

import (
	"fmt"
	"reflect"
	"testing"
)

func TestWithHTTPSetting_1(t *testing.T) {
	type args struct {
		setting BeegoHTTPSettings
	}
	tests := []struct {
		name string
		args args
		want ClientOption
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			if got := WithHTTPSetting(tt.args.setting); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("WithHTTPSetting() = %v, want %v", got, tt.want)
			}
		})
	}
}
