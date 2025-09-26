package port

import (
	"fmt"
	"reflect"
	"testing"
)

func TestWithRangePorts_1(t *testing.T) {
	type args struct {
		from int
		to   int
	}
	tests := []struct {
		name string
		args args
		want NameOption
	}{
		{
			name: "Test case 1",
			args: args{
				from: 1000,
				to:   2000,
			},
			want: func(builder *nameBuilder) *nameBuilder {
				builder.rangePortFrom = 1000
				builder.rangePortTo = 2000
				return builder
			},
		},
		{
			name: "Test case 2",
			args: args{
				from: 3000,
				to:   4000,
			},
			want: func(builder *nameBuilder) *nameBuilder {
				builder.rangePortFrom = 3000
				builder.rangePortTo = 4000
				return builder
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

			if got := WithRangePorts(tt.args.from, tt.args.to); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("WithRangePorts() = %v, want %v", got, tt.want)
			}
		})
	}
}
