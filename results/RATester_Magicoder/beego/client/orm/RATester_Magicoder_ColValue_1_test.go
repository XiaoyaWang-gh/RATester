package orm

import (
	"fmt"
	"reflect"
	"testing"
)

func TestColValue_1(t *testing.T) {
	type args struct {
		opt   operator
		value interface{}
	}
	tests := []struct {
		name string
		args args
		want interface{}
	}{
		{
			name: "TestColValue_0",
			args: args{
				opt:   ColAdd,
				value: "1",
			},
			want: colValue{
				value: 1,
				opt:   ColAdd,
			},
		},
		{
			name: "TestColValue_1",
			args: args{
				opt:   ColAdd,
				value: "a",
			},
			want: colValue{
				value: 0,
				opt:   ColAdd,
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

			if got := ColValue(tt.args.opt, tt.args.value); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ColValue() = %v, want %v", got, tt.want)
			}
		})
	}
}
