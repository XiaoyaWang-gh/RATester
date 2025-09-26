package hugo

import (
	"fmt"
	"reflect"
	"testing"
)

func TestNewInfo_1(t *testing.T) {
	type args struct {
		conf ConfigProvider
		deps []*Dependency
	}
	tests := []struct {
		name string
		args args
		want HugoInfo
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

			if got := NewInfo(tt.args.conf, tt.args.deps); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewInfo() = %v, want %v", got, tt.want)
			}
		})
	}
}
