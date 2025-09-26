package config

import (
	"fmt"
	"reflect"
	"testing"
)

func TestLoadFileContentWithTemplate_1(t *testing.T) {
	type args struct {
		path   string
		values *Values
	}
	tests := []struct {
		name    string
		args    args
		want    []byte
		wantErr bool
	}{
		{
			name: "test case 1",
			args: args{
				path:   "testdata/test.tmpl",
				values: &Values{Envs: map[string]string{"name": "test"}},
			},
			want:    []byte("test"),
			wantErr: false,
		},
		{
			name: "test case 2",
			args: args{
				path:   "testdata/test.tmpl",
				values: &Values{Envs: map[string]string{"name": "test"}},
			},
			want:    []byte("test"),
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			got, err := LoadFileContentWithTemplate(tt.args.path, tt.args.values)
			if (err != nil) != tt.wantErr {
				t.Errorf("LoadFileContentWithTemplate() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("LoadFileContentWithTemplate() = %v, want %v", got, tt.want)
			}
		})
	}
}
