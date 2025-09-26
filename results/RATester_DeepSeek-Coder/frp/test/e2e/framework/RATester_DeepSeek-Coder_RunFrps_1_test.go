package framework

import (
	"fmt"
	"testing"
)

func TestRunFrps_1(t *testing.T) {
	type args struct {
		args []string
	}
	tests := []struct {
		name    string
		f       *Framework
		args    args
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

			_, _, err := tt.f.RunFrps(tt.args.args...)
			if (err != nil) != tt.wantErr {
				t.Errorf("Framework.RunFrps() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
