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
		f         *Framework
		args      args
		wantOuts  []string
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

			gotOuts, gotPorts, err := tt.f.RenderTemplates(tt.args.templates)
			if (err != nil) != tt.wantErr {
				t.Errorf("Framework.RenderTemplates() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotOuts, tt.wantOuts) {
				t.Errorf("Framework.RenderTemplates() gotOuts = %v, want %v", gotOuts, tt.wantOuts)
			}
			if !reflect.DeepEqual(gotPorts, tt.wantPorts) {
				t.Errorf("Framework.RenderTemplates() gotPorts = %v, want %v", gotPorts, tt.wantPorts)
			}
		})
	}
}
