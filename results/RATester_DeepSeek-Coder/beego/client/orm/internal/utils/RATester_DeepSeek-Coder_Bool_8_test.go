package utils

import (
	"fmt"
	"testing"
)

func TestBool_8(t *testing.T) {
	tests := []struct {
		name    string
		f       StrTo
		want    bool
		wantErr bool
	}{
		{
			name: "TestBool_True",
			f:    StrTo("true"),
			want: true,
		},
		{
			name: "TestBool_False",
			f:    StrTo("false"),
			want: false,
		},
		{
			name:    "TestBool_Invalid",
			f:       StrTo("invalid"),
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

			got, err := tt.f.Bool()
			if (err != nil) != tt.wantErr {
				t.Errorf("StrTo.Bool() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("StrTo.Bool() = %v, want %v", got, tt.want)
			}
		})
	}
}
