package try

import (
	"fmt"
	"reflect"
	"testing"
)

func TestHasHeader_1(t *testing.T) {
	type args struct {
		header string
	}
	tests := []struct {
		name string
		args args
		want ResponseCondition
	}{
		{
			name: "TestHasHeader",
			args: args{header: "Content-Type"},
			want: HasHeader("Content-Type"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			if got := HasHeader(tt.args.header); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("HasHeader() = %v, want %v", got, tt.want)
			}
		})
	}
}
