package legacy

import (
	"fmt"
	"reflect"
	"testing"

	v1 "github.com/fatedier/frp/pkg/config/v1"
)

func TestConvert_ClientCommonConf_To_v1_1(t *testing.T) {
	tests := []struct {
		name string
		conf *ClientCommonConf
		want *v1.ClientCommonConfig
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

			got := Convert_ClientCommonConf_To_v1(tt.conf)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Convert_ClientCommonConf_To_v1() = %v, want %v", got, tt.want)
			}
		})
	}
}
