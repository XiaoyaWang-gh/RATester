package allconfig

import (
	"fmt"
	"testing"
)

func TestApplyOsEnvOverrides_1(t *testing.T) {
	type args struct {
		environ []string
	}
	tests := []struct {
		name    string
		l       configLoader
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

			if err := tt.l.applyOsEnvOverrides(tt.args.environ); (err != nil) != tt.wantErr {
				t.Errorf("configLoader.applyOsEnvOverrides() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
