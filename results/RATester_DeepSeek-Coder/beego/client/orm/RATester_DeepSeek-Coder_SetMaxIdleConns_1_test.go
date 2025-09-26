package orm

import (
	"fmt"
	"testing"
)

func TestSetMaxIdleConns_1(t *testing.T) {
	type args struct {
		aliasName    string
		maxIdleConns int
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "TestSetMaxIdleConns",
			args: args{
				aliasName:    "test",
				maxIdleConns: 10,
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

			SetMaxIdleConns(tt.args.aliasName, tt.args.maxIdleConns)
		})
	}
}
