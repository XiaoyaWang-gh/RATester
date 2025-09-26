package fiber

import (
	"fmt"
	"testing"
)

func TestExecuteOnRouteHooks_1(t *testing.T) {
	type args struct {
		route Route
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

			if err := tt.h.executeOnRouteHooks(tt.args.route); (err != nil) != tt.wantErr {
				t.Errorf("Hooks.executeOnRouteHooks() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
