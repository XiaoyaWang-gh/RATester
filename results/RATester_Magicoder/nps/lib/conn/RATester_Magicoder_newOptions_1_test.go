package conn

import (
	"fmt"
	"reflect"
	"testing"
	"time"
)

func TestnewOptions_1(t *testing.T) {
	type args struct {
		opts []Option
	}
	tests := []struct {
		name string
		args args
		want Options
	}{
		{
			name: "Test with default timeout",
			args: args{
				opts: []Option{},
			},
			want: Options{
				Timeout: defaultTimeOut,
			},
		},
		{
			name: "Test with custom timeout",
			args: args{
				opts: []Option{
					func(o *Options) {
						o.Timeout = 10 * time.Second
					},
				},
			},
			want: Options{
				Timeout: 10 * time.Second,
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

			if got := newOptions(tt.args.opts...); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("newOptions() = %v, want %v", got, tt.want)
			}
		})
	}
}
