package legacy

import (
	"fmt"
	"reflect"
	"testing"

	v1 "github.com/fatedier/frp/pkg/config/v1"
)

func TestConvert_ServerCommonConf_To_v1_1(t *testing.T) {
	type args struct {
		conf *ServerCommonConf
	}
	tests := []struct {
		name string
		args args
		want *v1.ServerConfig
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

			if got := Convert_ServerCommonConf_To_v1(tt.args.conf); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Convert_ServerCommonConf_To_v1() = %v, want %v", got, tt.want)
			}
		})
	}
}
