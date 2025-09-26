package web

import (
	"fmt"
	"reflect"
	"testing"
)

func TestDelStaticPath_1(t *testing.T) {
	tests := []struct {
		name string
		url  string
		want *HttpServer
	}{
		{
			name: "Test case 1",
			url:  "test",
			want: &HttpServer{},
		},
		{
			name: "Test case 2",
			url:  "/test",
			want: &HttpServer{},
		},
		{
			name: "Test case 3",
			url:  "test/",
			want: &HttpServer{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			if got := DelStaticPath(tt.url); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("DelStaticPath() = %v, want %v", got, tt.want)
			}
		})
	}
}
