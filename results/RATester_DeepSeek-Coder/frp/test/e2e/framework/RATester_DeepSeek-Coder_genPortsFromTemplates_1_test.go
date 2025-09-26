package framework

import (
	"fmt"
	"reflect"
	"testing"
)

func TestGenPortsFromTemplates_1(t *testing.T) {
	type args struct {
		templates []string
	}
	tests := []struct {
		name      string
		f         *Framework
		args      args
		wantPorts map[string]int
		wantErr   bool
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

			gotPorts, err := tt.f.genPortsFromTemplates(tt.args.templates)
			if (err != nil) != tt.wantErr {
				t.Errorf("Framework.genPortsFromTemplates() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotPorts, tt.wantPorts) {
				t.Errorf("Framework.genPortsFromTemplates() = %v, want %v", gotPorts, tt.wantPorts)
			}
		})
	}
}
