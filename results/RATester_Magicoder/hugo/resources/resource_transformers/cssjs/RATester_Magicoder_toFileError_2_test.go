package cssjs

import (
	"fmt"
	"testing"
)

func TesttoFileError_2(t *testing.T) {
	type args struct {
		output string
	}
	tests := []struct {
		name    string
		imp     *importResolver
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

			if err := tt.imp.toFileError(tt.args.output); (err != nil) != tt.wantErr {
				t.Errorf("importResolver.toFileError() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
