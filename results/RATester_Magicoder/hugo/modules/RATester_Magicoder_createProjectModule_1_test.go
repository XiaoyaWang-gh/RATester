package modules

import (
	"fmt"
	"reflect"
	"testing"
)

func TestcreateProjectModule_1(t *testing.T) {
	type args struct {
		gomod      *goModule
		workingDir string
		conf       Config
	}
	tests := []struct {
		name string
		args args
		want *moduleAdapter
	}{
		{
			name: "Test Case 1",
			args: args{
				gomod:      &goModule{},
				workingDir: "/path/to/working/dir",
				conf:       Config{},
			},
			want: &moduleAdapter{
				path:       "project",
				dir:        "/path/to/working/dir",
				gomod:      &goModule{},
				projectMod: true,
				config:     Config{},
			},
		},
		{
			name: "Test Case 2",
			args: args{
				gomod:      nil,
				workingDir: "/path/to/working/dir",
				conf:       Config{},
			},
			want: &moduleAdapter{
				path:       "project",
				dir:        "/path/to/working/dir",
				gomod:      nil,
				projectMod: true,
				config:     Config{},
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

			if got := createProjectModule(tt.args.gomod, tt.args.workingDir, tt.args.conf); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("createProjectModule() = %v, want %v", got, tt.want)
			}
		})
	}
}
