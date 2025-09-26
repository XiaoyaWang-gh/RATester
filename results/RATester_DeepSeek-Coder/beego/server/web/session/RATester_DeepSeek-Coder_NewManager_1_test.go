package session

import (
	"fmt"
	"reflect"
	"testing"
)

func TestNewManager_1(t *testing.T) {
	type args struct {
		provideName string
		cf          *ManagerConfig
	}
	tests := []struct {
		name    string
		args    args
		want    *Manager
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

			got, err := NewManager(tt.args.provideName, tt.args.cf)
			if (err != nil) != tt.wantErr {
				t.Errorf("NewManager() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewManager() = %v, want %v", got, tt.want)
			}
		})
	}
}
