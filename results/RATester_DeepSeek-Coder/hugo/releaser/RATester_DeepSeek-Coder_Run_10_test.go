package releaser

import (
	"fmt"
	"testing"
)

func TestRun_10(t *testing.T) {
	type fields struct {
		branchVersion string
		step          int
		skipPush      bool
		try           bool
		git           func(args ...string) (string, error)
	}
	tests := []struct {
		name    string
		fields  fields
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			r := &ReleaseHandler{
				branchVersion: tt.fields.branchVersion,
				step:          tt.fields.step,
				skipPush:      tt.fields.skipPush,
				try:           tt.fields.try,
				git:           tt.fields.git,
			}
			if err := r.Run(); (err != nil) != tt.wantErr {
				t.Errorf("ReleaseHandler.Run() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
