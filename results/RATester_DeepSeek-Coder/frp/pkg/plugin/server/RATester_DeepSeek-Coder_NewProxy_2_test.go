package plugin

import (
	"fmt"
	"reflect"
	"testing"
)

func TestNewProxy_2(t *testing.T) {
	type args struct {
		content *NewProxyContent
	}
	tests := []struct {
		name    string
		m       *Manager
		args    args
		want    *NewProxyContent
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

			got, err := tt.m.NewProxy(tt.args.content)
			if (err != nil) != tt.wantErr {
				t.Errorf("Manager.NewProxy() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Manager.NewProxy() = %v, want %v", got, tt.want)
			}
		})
	}
}
