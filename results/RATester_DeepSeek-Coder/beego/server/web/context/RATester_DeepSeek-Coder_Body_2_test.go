package context

import (
	"fmt"
	"testing"
)

func TestBody_2(t *testing.T) {
	type args struct {
		content []byte
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
			if err := output.Body(tt.args.content); (err != nil) != tt.wantErr {
				t.Errorf("BeegoOutput.Body() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
