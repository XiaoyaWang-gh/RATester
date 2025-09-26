package template

import (
	"fmt"
	"reflect"
	"testing"
)

func TestParse_7(t *testing.T) {
	type args struct {
		text string
	}
	tests := []struct {
		name    string
		t       *Template
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

			got, err := tt.t.Parse(tt.args.text)
			if (err != nil) != tt.wantErr {
				t.Errorf("Template.Parse() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Template.Parse() = %v, want %v", got, tt.want)
			}
		})
	}
}
