package legacy

import (
	"fmt"
	"reflect"
	"testing"
)

func TestNewConfByType_1(t *testing.T) {
	type args struct {
		proxyType ProxyType
	}
	tests := []struct {
		name string
		args args
		want ProxyConf
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

			if got := NewConfByType(tt.args.proxyType); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewConfByType() = %v, want %v", got, tt.want)
			}
		})
	}
}
