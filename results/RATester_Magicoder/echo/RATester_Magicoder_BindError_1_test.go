package echo

import (
	"fmt"
	"testing"
)

func TestBindError_1(t *testing.T) {
	tests := []struct {
		name    string
		b       *ValueBinder
		wantErr bool
	}{
		{
			name: "Test case 1",
			b: &ValueBinder{
				errors: []error{fmt.Errorf("error 1")},
			},
			wantErr: true,
		},
		{
			name: "Test case 2",
			b: &ValueBinder{
				errors: nil,
			},
			wantErr: false,
		},
		{
			name: "Test case 3",
			b: &ValueBinder{
				errors: []error{fmt.Errorf("error 1"), fmt.Errorf("error 2")},
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

			if err := tt.b.BindError(); (err != nil) != tt.wantErr {
				t.Errorf("ValueBinder.BindError() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
