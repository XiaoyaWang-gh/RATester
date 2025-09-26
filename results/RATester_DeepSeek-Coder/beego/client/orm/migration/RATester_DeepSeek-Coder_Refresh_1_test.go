package migration

import (
	"fmt"
	"testing"
)

func TestRefresh_1(t *testing.T) {
	tests := []struct {
		name    string
		wantErr bool
	}{
		{"Test Case 1", false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			if err := Refresh(); (err != nil) != tt.wantErr {
				t.Errorf("Refresh() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
