package mock

import (
	"fmt"
	"testing"
)

func TestRegisterMockDB_1(t *testing.T) {
	type args struct {
		name string
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "Test case 1",
			args: args{
				name: "test",
			},
		},
		{
			name: "Test case 2",
			args: args{
				name: "test2",
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

			RegisterMockDB(tt.args.name)
		})
	}
}
