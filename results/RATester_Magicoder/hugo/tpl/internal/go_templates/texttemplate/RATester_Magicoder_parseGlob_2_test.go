package template

import (
	"fmt"
	"reflect"
	"testing"
)

func TestparseGlob_2(t *testing.T) {
	type args struct {
		t       *Template
		pattern string
	}
	tests := []struct {
		name    string
		args    args
		want    *Template
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

			got, err := parseGlob(tt.args.t, tt.args.pattern)
			if (err != nil) != tt.wantErr {
				t.Errorf("parseGlob() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("parseGlob() = %v, want %v", got, tt.want)
			}
		})
	}
}
