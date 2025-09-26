package web

import (
	"fmt"
	"reflect"
	"testing"
)

func TestSetStaticPath_1(t *testing.T) {
	type args struct {
		url  string
		path string
	}
	tests := []struct {
		name string
		args args
		want *HttpServer
	}{
		{
			name: "Test case 1",
			args: args{
				url:  "/static",
				path: "/path/to/static",
			},
			want: BeeApp,
		},
		{
			name: "Test case 2",
			args: args{
				url:  "static",
				path: "/path/to/static",
			},
			want: BeeApp,
		},
		{
			name: "Test case 3",
			args: args{
				url:  "/static/",
				path: "/path/to/static",
			},
			want: BeeApp,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			if got := SetStaticPath(tt.args.url, tt.args.path); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SetStaticPath() = %v, want %v", got, tt.want)
			}
		})
	}
}
