package log

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/beego/beego/v2/core/logs"
)

func TestNewFilterChainBuilder_1(t *testing.T) {
	type args struct {
		opts []BuilderOption
	}
	tests := []struct {
		name string
		args args
		want *FilterChainBuilder
	}{
		{
			name: "test1",
			args: args{
				opts: []BuilderOption{
					func(builder *FilterChainBuilder) {
						builder.printableContentTypes = []string{"application/json"}
					},
				},
			},
			want: &FilterChainBuilder{
				printableContentTypes: []string{"application/json"},
				log:                   logs.Debug,
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

			if got := NewFilterChainBuilder(tt.args.opts...); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewFilterChainBuilder() = %v, want %v", got, tt.want)
			}
		})
	}
}
