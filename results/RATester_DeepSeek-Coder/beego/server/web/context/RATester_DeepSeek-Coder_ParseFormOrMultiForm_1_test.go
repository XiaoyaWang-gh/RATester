package context

import (
	"fmt"
	"testing"
)

func TestParseFormOrMultiForm_1(t *testing.T) {
	type args struct {
		maxMemory int64
	}
	tests := []struct {
		name    string
		input   *BeegoInput
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

			if err := tt.input.ParseFormOrMultiForm(tt.args.maxMemory); (err != nil) != tt.wantErr {
				t.Errorf("BeegoInput.ParseFormOrMultiForm() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
