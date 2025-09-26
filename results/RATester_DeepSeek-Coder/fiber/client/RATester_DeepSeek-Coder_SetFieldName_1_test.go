package client

import (
	"fmt"
	"testing"
)

func TestSetFieldName_1(t *testing.T) {
	type args struct {
		n string
	}
	tests := []struct {
		name string
		f    *File
		args args
		want string
	}{
		{
			name: "TestSetFieldName",
			f:    &File{},
			args: args{n: "testFieldName"},
			want: "testFieldName",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			tt.f.SetFieldName(tt.args.n)
			if got := tt.f.fieldName; got != tt.want {
				t.Errorf("File.SetFieldName() = %v, want %v", got, tt.want)
			}
		})
	}
}
