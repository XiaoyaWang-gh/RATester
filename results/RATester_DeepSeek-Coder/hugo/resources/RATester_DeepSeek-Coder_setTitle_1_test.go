package resources

import (
	"fmt"
	"testing"
)

func TestSetTitle_1(t *testing.T) {
	type args struct {
		title string
	}
	tests := []struct {
		name string
		r    *metaResource
		args args
		want string
	}{
		{
			name: "Test case 1",
			r:    &metaResource{},
			args: args{title: "Test Title"},
			want: "Test Title",
		},
		{
			name: "Test case 2",
			r:    &metaResource{},
			args: args{title: "Another Title"},
			want: "Another Title",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			tt.r.setTitle(tt.args.title)
			if got := tt.r.Title(); got != tt.want {
				t.Errorf("metaResource.setTitle() = %v, want %v", got, tt.want)
			}
		})
	}
}
