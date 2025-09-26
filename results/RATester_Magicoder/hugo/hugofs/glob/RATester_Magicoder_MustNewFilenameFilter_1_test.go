package glob

import (
	"fmt"
	"reflect"
	"strings"
	"testing"
)

func TestMustNewFilenameFilter_1(t *testing.T) {
	type args struct {
		inclusions []string
		exclusions []string
	}
	tests := []struct {
		name string
		args args
		want *FilenameFilter
	}{
		{
			name: "Test case 1",
			args: args{
				inclusions: []string{"*.txt"},
				exclusions: []string{},
			},
			want: &FilenameFilter{
				shouldInclude: func(filename string) bool {
					return strings.HasSuffix(filename, ".txt")
				},
			},
		},
		{
			name: "Test case 2",
			args: args{
				inclusions: []string{"*.go"},
				exclusions: []string{"*_test.go"},
			},
			want: &FilenameFilter{
				shouldInclude: func(filename string) bool {
					return strings.HasSuffix(filename, ".go") && !strings.HasSuffix(filename, "_test.go")
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

			if got := MustNewFilenameFilter(tt.args.inclusions, tt.args.exclusions); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("MustNewFilenameFilter() = %v, want %v", got, tt.want)
			}
		})
	}
}
