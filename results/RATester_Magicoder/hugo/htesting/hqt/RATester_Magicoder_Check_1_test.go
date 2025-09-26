package hqt

import (
	"fmt"
	"testing"
)

func TestCheck_1(t *testing.T) {
	type args struct {
		got  any
		args []any
		note func(key string, value any)
	}
	tests := []struct {
		name    string
		c       *stringChecker
		args    args
		wantErr bool
	}{
		{
			name: "Test case 1",
			c:    &stringChecker{},
			args: args{
				got:  "Hello, World!",
				args: []any{"Hello, World!"},
				note: func(key string, value any) {},
			},
			wantErr: false,
		},
		{
			name: "Test case 2",
			c:    &stringChecker{},
			args: args{
				got:  "Hello, World!",
				args: []any{"Hello, Universe!"},
				note: func(key string, value any) {},
			},
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

			if err := tt.c.Check(tt.args.got, tt.args.args, tt.args.note); (err != nil) != tt.wantErr {
				t.Errorf("stringChecker.Check() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
