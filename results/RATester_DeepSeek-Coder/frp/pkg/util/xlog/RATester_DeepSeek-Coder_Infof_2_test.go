package xlog

import (
	"fmt"
	"testing"
)

func TestInfof_2(t *testing.T) {
	type args struct {
		format string
		v      []interface{}
	}
	tests := []struct {
		name string
		l    *Logger
		args args
		want string
	}{
		{
			name: "TestInfof",
			l:    &Logger{prefixString: "test_prefix"},
			args: args{
				format: "test_format",
				v:      []interface{}{"test_arg"},
			},
			want: "test_prefixtest_format: test_arg",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			tt.l.Infof(tt.args.format, tt.args.v...)
			if tt.l.prefixString != tt.want {
				t.Errorf("Logger.Infof() = %v, want %v", tt.l.prefixString, tt.want)
			}
		})
	}
}
