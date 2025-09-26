package hexec

import (
	"fmt"
	"reflect"
	"testing"
)

func Testcommand_1(t *testing.T) {
	type args struct {
		arg []any
	}
	tests := []struct {
		name    string
		c       *commandeer
		args    args
		want    *cmdWrapper
		wantErr bool
	}{
		{
			name: "test case 1",
			c: &commandeer{
				name: "test",
			},
			args: args{
				arg: []any{"arg1", "arg2"},
			},
			want: &cmdWrapper{
				name: "test",
			},
			wantErr: false,
		},
		{
			name: "test case 2",
			c: &commandeer{
				name: "test",
			},
			args: args{
				arg: []any{"arg1", func(c *commandeer) {
					c.name = "new name"
				}},
			},
			want: &cmdWrapper{
				name: "new name",
			},
			wantErr: false,
		},
		{
			name: "test case 3",
			c: &commandeer{
				name: "test",
			},
			args: args{
				arg: []any{"arg1", 123},
			},
			want:    nil,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			got, err := tt.c.command(tt.args.arg...)
			if (err != nil) != tt.wantErr {
				t.Errorf("command() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("command() = %v, want %v", got, tt.want)
			}
		})
	}
}
