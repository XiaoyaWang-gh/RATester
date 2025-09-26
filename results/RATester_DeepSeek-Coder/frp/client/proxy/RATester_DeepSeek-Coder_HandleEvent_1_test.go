package proxy

import (
	"fmt"
	"testing"
)

func TestHandleEvent_1(t *testing.T) {
	type args struct {
		payload interface{}
	}
	tests := []struct {
		name    string
		pm      *Manager
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

			err := tt.pm.HandleEvent(tt.args.payload)
			if (err != nil) != tt.wantErr {
				t.Errorf("Manager.HandleEvent() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
