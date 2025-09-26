package fiber

import (
	"fmt"
	"testing"
)

func TestExecuteOnGroupNameHooks_1(t *testing.T) {
	type args struct {
		group Group
	}
	tests := []struct {
		name    string
		h       *Hooks
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

			if err := tt.h.executeOnGroupNameHooks(tt.args.group); (err != nil) != tt.wantErr {
				t.Errorf("Hooks.executeOnGroupNameHooks() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
