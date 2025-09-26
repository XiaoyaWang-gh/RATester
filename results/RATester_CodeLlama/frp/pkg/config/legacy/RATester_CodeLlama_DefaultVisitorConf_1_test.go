package legacy

import (
	"fmt"
	"reflect"
	"testing"
)

func TestDefaultVisitorConf_1(t *testing.T) {
	type args struct {
		visitorType VisitorType
	}
	tests := []struct {
		name string
		args args
		want VisitorConf
	}{
		{
			name: "test case 1",
			args: args{
				visitorType: VisitorType(""),
			},
			want: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			if got := DefaultVisitorConf(tt.args.visitorType); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("DefaultVisitorConf() = %v, want %v", got, tt.want)
			}
		})
	}
}
