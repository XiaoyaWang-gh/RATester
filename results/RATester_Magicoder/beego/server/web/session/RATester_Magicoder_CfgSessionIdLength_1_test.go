package session

import (
	"fmt"
	"reflect"
	"testing"
)

func TestCfgSessionIdLength_1(t *testing.T) {
	type args struct {
		length int64
	}
	tests := []struct {
		name string
		args args
		want ManagerConfigOpt
	}{
		{
			name: "TestCfgSessionIdLength",
			args: args{length: 10},
			want: func(config *ManagerConfig) {
				config.SessionIDLength = 10
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			if got := CfgSessionIdLength(tt.args.length); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("CfgSessionIdLength() = %v, want %v", got, tt.want)
			}
		})
	}
}
