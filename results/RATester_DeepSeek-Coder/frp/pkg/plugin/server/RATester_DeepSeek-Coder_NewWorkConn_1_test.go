package plugin

import (
	"fmt"
	"reflect"
	"testing"
)

func TestNewWorkConn_1(t *testing.T) {
	type args struct {
		content *NewWorkConnContent
	}
	tests := []struct {
		name    string
		m       *Manager
		args    args
		want    *NewWorkConnContent
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

			got, err := tt.m.NewWorkConn(tt.args.content)
			if (err != nil) != tt.wantErr {
				t.Errorf("Manager.NewWorkConn() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Manager.NewWorkConn() = %v, want %v", got, tt.want)
			}
		})
	}
}
