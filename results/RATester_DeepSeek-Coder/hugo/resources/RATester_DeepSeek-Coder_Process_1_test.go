package resources

import (
	"fmt"
	"testing"
)

func TestProcess_1(t *testing.T) {
	tests := []struct {
		name    string
		spec    string
		wantErr bool
	}{
		{
			name:    "valid spec",
			spec:    "fill 100x100",
			wantErr: false,
		},
		{
			name:    "invalid spec",
			spec:    "invalid",
			wantErr: true,
		},
		{
			name:    "empty spec",
			spec:    "",
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			i := &imageResource{}
			_, err := i.Process(tt.spec)
			if (err != nil) != tt.wantErr {
				t.Errorf("imageResource.Process() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
