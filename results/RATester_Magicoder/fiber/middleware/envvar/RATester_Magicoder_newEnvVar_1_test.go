package envvar

import (
	"fmt"
	"reflect"
	"testing"
)

func TestnewEnvVar_1(t *testing.T) {
	type args struct {
		cfg Config
	}
	tests := []struct {
		name string
		args args
		want *EnvVar
	}{
		{
			name: "Test Case 1",
			args: args{
				cfg: Config{
					ExportVars: map[string]string{
						"key1": "value1",
						"key2": "value2",
					},
					ExcludeVars: map[string]string{
						"key3": "value3",
						"key4": "value4",
					},
				},
			},
			want: &EnvVar{
				Vars: map[string]string{
					"key1": "value1",
					"key2": "value2",
				},
			},
		},
		{
			name: "Test Case 2",
			args: args{
				cfg: Config{
					ExportVars: map[string]string{
						"key5": "value5",
						"key6": "value6",
					},
					ExcludeVars: map[string]string{
						"key7": "value7",
						"key8": "value8",
					},
				},
			},
			want: &EnvVar{
				Vars: map[string]string{
					"key5": "value5",
					"key6": "value6",
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

			if got := newEnvVar(tt.args.cfg); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("newEnvVar() = %v, want %v", got, tt.want)
			}
		})
	}
}
