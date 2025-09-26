package orm

import (
	"fmt"
	"testing"
)

func TestSetMaxOpenConns_2(t *testing.T) {
	type args struct {
		aliasName    string
		maxOpenConns int
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "TestSetMaxOpenConns",
			args: args{
				aliasName:    "test",
				maxOpenConns: 10,
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

			SetMaxOpenConns(tt.args.aliasName, tt.args.maxOpenConns)
		})
	}
}
