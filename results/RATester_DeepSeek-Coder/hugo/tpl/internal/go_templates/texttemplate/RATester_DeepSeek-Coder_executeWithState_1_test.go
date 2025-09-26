package template

import (
	"fmt"
	"reflect"
	"testing"
)

func TestExecuteWithState_1(t *testing.T) {
	type args struct {
		state *state
		value reflect.Value
	}
	tests := []struct {
		name    string
		t       *Template
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

			if err := tt.t.executeWithState(tt.args.state, tt.args.value); (err != nil) != tt.wantErr {
				t.Errorf("Template.executeWithState() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
