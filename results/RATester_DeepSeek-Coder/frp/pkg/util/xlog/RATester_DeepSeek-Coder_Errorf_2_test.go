package xlog

import (
	"fmt"
	"testing"
)

func TestErrorf_2(t *testing.T) {
	type fields struct {
		prefixes     []LogPrefix
		prefixString string
	}
	type args struct {
		format string
		v      []interface{}
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   string
	}{
		{
			name: "TestErrorf",
			fields: fields{
				prefixes: []LogPrefix{
					{
						Name:     "test",
						Value:    "test",
						Priority: 10,
					},
				},
				prefixString: "[test:test] ",
			},
			args: args{
				format: "test %s",
				v:      []interface{}{"value"},
			},
			want: "[test:test] test value",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			l := &Logger{
				prefixes:     tt.fields.prefixes,
				prefixString: tt.fields.prefixString,
			}
			l.Errorf(tt.args.format, tt.args.v...)
			// TODO: Add assertion to check the log output
		})
	}
}
