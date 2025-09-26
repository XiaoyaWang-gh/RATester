package web

import (
	"fmt"
	"testing"
)

func TestNotNil_1(t *testing.T) {
	type args struct {
		a interface{}
	}
	tests := []struct {
		name      string
		args      args
		wantIsNil bool
	}{
		{
			name: "Test case 1",
			args: args{
				a: "test",
			},
			wantIsNil: false,
		},
		{
			name: "Test case 2",
			args: args{
				a: nil,
			},
			wantIsNil: true,
		},
		{
			name: "Test case 3",
			args: args{
				a: 0,
			},
			wantIsNil: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			if gotIsNil := NotNil(tt.args.a); gotIsNil != tt.wantIsNil {
				t.Errorf("NotNil() = %v, want %v", gotIsNil, tt.wantIsNil)
			}
		})
	}
}
