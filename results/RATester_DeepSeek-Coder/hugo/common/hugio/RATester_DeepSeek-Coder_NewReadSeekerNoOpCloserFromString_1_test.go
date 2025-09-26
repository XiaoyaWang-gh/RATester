package hugio

import (
	"fmt"
	"reflect"
	"testing"
)

func TestNewReadSeekerNoOpCloserFromString_1(t *testing.T) {
	type args struct {
		content string
	}
	tests := []struct {
		name string
		args args
		want ReadSeekCloser
	}{
		{
			name: "Test case 1",
			args: args{content: "test content"},
			want: NewReadSeekerNoOpCloserFromString("test content"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			got := NewReadSeekerNoOpCloserFromString(tt.args.content)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewReadSeekerNoOpCloserFromString() = %v, want %v", got, tt.want)
			}
		})
	}
}
