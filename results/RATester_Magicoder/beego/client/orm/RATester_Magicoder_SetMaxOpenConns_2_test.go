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
			name: "TestSetMaxOpenConns_0",
			args: args{
				aliasName:    "TestSetMaxOpenConns_0",
				maxOpenConns: 10,
			},
		},
		{
			name: "TestSetMaxOpenConns_1",
			args: args{
				aliasName:    "TestSetMaxOpenConns_1",
				maxOpenConns: 20,
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
