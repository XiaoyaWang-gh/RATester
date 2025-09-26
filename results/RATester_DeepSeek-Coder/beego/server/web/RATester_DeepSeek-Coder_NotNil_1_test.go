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
			name: "Test with nil",
			args: args{
				a: nil,
			},
			wantIsNil: true,
		},
		{
			name: "Test with non-nil string",
			args: args{
				a: "test",
			},
			wantIsNil: false,
		},
		{
			name: "Test with non-nil int",
			args: args{
				a: 1,
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
