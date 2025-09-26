package commands

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/bep/simplecobra"
)

func TestNewExec_1(t *testing.T) {
	tests := []struct {
		name    string
		want    *simplecobra.Exec
		wantErr bool
	}{
		{
			name:    "Test Case 1",
			want:    &simplecobra.Exec{},
			wantErr: false,
		},
		// Add more test cases as needed
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			got, err := newExec()
			if (err != nil) != tt.wantErr {
				t.Errorf("newExec() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("newExec() = %v, want %v", got, tt.want)
			}
		})
	}
}
