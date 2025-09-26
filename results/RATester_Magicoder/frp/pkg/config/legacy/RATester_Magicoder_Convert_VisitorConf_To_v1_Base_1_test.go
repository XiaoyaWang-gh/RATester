package legacy

import (
	"fmt"
	"reflect"
	"testing"

	v1 "github.com/fatedier/frp/pkg/config/v1"
)

func TestConvert_VisitorConf_To_v1_Base_1(t *testing.T) {
	type args struct {
		conf VisitorConf
	}
	tests := []struct {
		name string
		args args
		want *v1.VisitorBaseConfig
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

			if got := Convert_VisitorConf_To_v1_Base(tt.args.conf); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Convert_VisitorConf_To_v1_Base() = %v, want %v", got, tt.want)
			}
		})
	}
}
