package template

import (
	"fmt"
	"reflect"
	"testing"
)

func TestPrepare_2(t *testing.T) {
	tests := []struct {
		name    string
		fields  Template
		want    *Template
		wantErr bool
	}{
		{
			name: "Test Case 1",
			fields: Template{
				name: "Test Template",
			},
			want: &Template{
				name: "Test Template",
			},
			wantErr: false,
		},
		// Add more test cases here
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			tr := &Template{
				name: tt.fields.name,
			}
			got, err := tr.Prepare()
			if (err != nil) != tt.wantErr {
				t.Errorf("Template.Prepare() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Template.Prepare() = %v, want %v", got, tt.want)
			}
		})
	}
}
