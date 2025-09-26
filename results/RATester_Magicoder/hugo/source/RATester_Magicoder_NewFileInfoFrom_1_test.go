package source

import (
	"fmt"
	"reflect"
	"testing"
)

func TestNewFileInfoFrom_1(t *testing.T) {
	type args struct {
		path     string
		filename string
	}
	tests := []struct {
		name string
		args args
		want *File
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

			if got := NewFileInfoFrom(tt.args.path, tt.args.filename); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewFileInfoFrom() = %v, want %v", got, tt.want)
			}
		})
	}
}
