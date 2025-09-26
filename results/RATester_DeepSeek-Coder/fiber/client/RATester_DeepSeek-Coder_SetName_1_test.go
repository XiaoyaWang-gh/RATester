package client

import (
	"fmt"
	"testing"
)

func TestSetName_1(t *testing.T) {
	type args struct {
		n string
	}
	tests := []struct {
		name string
		f    *File
		args args
		want string
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

			f := &File{
				reader:    tt.f.reader,
				name:      tt.f.name,
				fieldName: tt.f.fieldName,
				path:      tt.f.path,
			}
			f.SetName(tt.args.n)
			if f.name != tt.want {
				t.Errorf("File.SetName() = %v, want %v", f.name, tt.want)
			}
		})
	}
}
