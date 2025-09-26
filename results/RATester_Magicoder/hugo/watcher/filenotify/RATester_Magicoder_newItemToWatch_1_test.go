package filenotify

import (
	"fmt"
	"reflect"
	"testing"
)

func TestnewItemToWatch_1(t *testing.T) {
	type args struct {
		filename string
	}
	tests := []struct {
		name    string
		args    args
		want    *itemToWatch
		wantErr bool
	}{
		{
			name: "Test Case 1",
			args: args{
				filename: "test.txt",
			},
			want: &itemToWatch{
				filename: "test.txt",
				left:     &recording{},
			},
			wantErr: false,
		},
		{
			name: "Test Case 2",
			args: args{
				filename: "test2.txt",
			},
			want: &itemToWatch{
				filename: "test2.txt",
				left:     &recording{},
			},
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

			got, err := newItemToWatch(tt.args.filename)
			if (err != nil) != tt.wantErr {
				t.Errorf("newItemToWatch() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("newItemToWatch() = %v, want %v", got, tt.want)
			}
		})
	}
}
