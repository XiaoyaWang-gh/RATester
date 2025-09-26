package visitor

import (
	"fmt"
	"testing"

	v1 "github.com/fatedier/frp/pkg/config/v1"
)

func TestStartVisitor_1(t *testing.T) {
	type args struct {
		cfg v1.VisitorConfigurer
	}
	tests := []struct {
		name    string
		vm      *Manager
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

			if err := tt.vm.startVisitor(tt.args.cfg); (err != nil) != tt.wantErr {
				t.Errorf("Manager.startVisitor() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
