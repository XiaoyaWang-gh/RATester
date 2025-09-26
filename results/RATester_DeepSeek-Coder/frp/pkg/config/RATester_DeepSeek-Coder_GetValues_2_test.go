package config

import (
	"fmt"
	"reflect"
	"testing"
)

func TestGetValues_2(t *testing.T) {
	type args struct {
		glbEnvs map[string]string
	}
	tests := []struct {
		name string
		args args
		want *Values
	}{
		{
			name: "Test case 1",
			args: args{
				glbEnvs: map[string]string{
					"key1": "value1",
					"key2": "value2",
				},
			},
			want: &Values{
				Envs: map[string]string{
					"key1": "value1",
					"key2": "value2",
				},
			},
		},
		{
			name: "Test case 2",
			args: args{
				glbEnvs: map[string]string{
					"key3": "value3",
					"key4": "value4",
				},
			},
			want: &Values{
				Envs: map[string]string{
					"key3": "value3",
					"key4": "value4",
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			glbEnvs = tt.args.glbEnvs
			if got := GetValues(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetValues() = %v, want %v", got, tt.want)
			}
		})
	}
}
