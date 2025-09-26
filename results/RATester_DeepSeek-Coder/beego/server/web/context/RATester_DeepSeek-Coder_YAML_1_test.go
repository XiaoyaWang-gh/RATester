package context

import (
	"fmt"
	"testing"
)

func TestYAML_1(t *testing.T) {
	type args struct {
		data interface{}
	}
	tests := []struct {
		name    string
		output  *BeegoOutput
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

			output := &BeegoOutput{}
			if err := output.YAML(tt.args.data); (err != nil) != tt.wantErr {
				t.Errorf("BeegoOutput.YAML() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
