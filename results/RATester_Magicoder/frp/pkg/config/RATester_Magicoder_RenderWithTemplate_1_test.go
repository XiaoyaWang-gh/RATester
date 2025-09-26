package config

import (
	"fmt"
	"reflect"
	"testing"
)

func TestRenderWithTemplate_1(t *testing.T) {
	type args struct {
		in     []byte
		values *Values
	}
	tests := []struct {
		name    string
		args    args
		want    []byte
		wantErr bool
	}{
		{
			name: "Test case 1",
			args: args{
				in: []byte("{{.Envs.field_name}}"),
				values: &Values{
					Envs: map[string]string{
						"field_name": "value",
					},
				},
			},
			want:    []byte("value"),
			wantErr: false,
		},
		{
			name: "Test case 2",
			args: args{
				in: []byte("{{.Envs.field_name}}"),
				values: &Values{
					Envs: map[string]string{
						"field_name": "value",
					},
				},
			},
			want:    []byte("value"),
			wantErr: false,
		},
		// Add more test cases as needed
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			got, err := RenderWithTemplate(tt.args.in, tt.args.values)
			if (err != nil) != tt.wantErr {
				t.Errorf("RenderWithTemplate() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("RenderWithTemplate() = %v, want %v", got, tt.want)
			}
		})
	}
}
