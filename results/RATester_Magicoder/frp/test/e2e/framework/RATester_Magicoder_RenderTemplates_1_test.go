package framework

import (
	"fmt"
	"reflect"
	"testing"
)

func TestRenderTemplates_1(t *testing.T) {
	type args struct {
		templates []string
	}
	tests := []struct {
		name      string
		args      args
		wantOut   []string
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

			f := &Framework{}
			gotOuts, gotPorts, err := f.RenderTemplates(tt.args.templates)
			if (err != nil) != tt.wantErr {
				t.Errorf("RenderTemplates() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotOuts, tt.wantOut) {
				t.Errorf("RenderTemplates() gotOuts = %v, want %v", gotOuts, tt.wantOut)
			}
			if !reflect.DeepEqual(gotPorts, tt.wantPorts) {
				t.Errorf("RenderTemplates() gotPorts = %v, want %v", gotPorts, tt.wantPorts)
			}
		})
	}
}
