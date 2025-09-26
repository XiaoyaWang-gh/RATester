package resources

import (
	"fmt"
	"testing"
)

func TestNewResource_1(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name    string
		rd      ResourceSourceDescriptor
		wantErr bool
	}{
		{
			name: "success",
			rd: ResourceSourceDescriptor{
				TargetPath: "/path/to/file",
			},
			wantErr: false,
		},
		{
			name: "failure",
			rd: ResourceSourceDescriptor{
				TargetPath: "/path/to/file",
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

			r := &Spec{}
			_, err := r.NewResource(tt.rd)
			if (err != nil) != tt.wantErr {
				t.Errorf("NewResource() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}
