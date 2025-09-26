package fiber

import (
	"fmt"
	"testing"
)

func TestExecuteOnListenHooks_1(t *testing.T) {
	type args struct {
		listenData ListenData
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

			if err := tt.h.executeOnListenHooks(tt.args.listenData); (err != nil) != tt.wantErr {
				t.Errorf("Hooks.executeOnListenHooks() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
