package conn

import (
	"fmt"
	"reflect"
	"testing"
	"time"
)

func TestNewOptions_1(t *testing.T) {
	type args struct {
		opts []Option
	}
	tests := []struct {
		name string
		args args
		want Options
	}{
		{
			name: "test_newOptions_001",
			args: args{
				opts: []Option{
					func(opt *Options) {
						opt.Timeout = 10 * time.Second
					},
				},
			},
			want: Options{
				Timeout: 10 * time.Second,
			},
		},
		{
			name: "test_newOptions_002",
			args: args{
				opts: []Option{
					func(opt *Options) {
						opt.Timeout = 10 * time.Second
					},
					func(opt *Options) {
						opt.Timeout = 20 * time.Second
					},
				},
			},
			want: Options{
				Timeout: 20 * time.Second,
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
