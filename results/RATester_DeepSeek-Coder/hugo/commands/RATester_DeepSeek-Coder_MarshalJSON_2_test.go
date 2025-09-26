package commands

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/gohugoio/hugo/modules"
)

func TestMarshalJSON_2(t *testing.T) {
	type fields struct {
		verbose bool
		m       modules.Module
	}
	tests := []struct {
		name    string
		fields  fields
		want    []byte
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

			m := &configModMounts{
				verbose: tt.fields.verbose,
				m:       tt.fields.m,
			}
			got, err := m.MarshalJSON()
			if (err != nil) != tt.wantErr {
				t.Errorf("configModMounts.MarshalJSON() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("configModMounts.MarshalJSON() = %v, want %v", got, tt.want)
			}
		})
	}
}
