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
		name    string
		args    args
		want    map[string]int
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

			f := &Framework{}
			got, err := f.genPortsFromTemplates(tt.args.templates)
			if (err != nil) != tt.wantErr {
				t.Errorf("genPortsFromTemplates() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("genPortsFromTemplates() = %v, want %v", got, tt.want)
			}
		})
	}
}
