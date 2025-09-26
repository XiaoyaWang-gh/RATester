package vhost

import (
	"fmt"
	"reflect"
	"testing"
)

func TestGetNotFoundPageContent_1(t *testing.T) {
	tests := []struct {
		name     string
		want     []byte
		setup    func()
		teardown func()
	}{
		{
			name: "Test case 1",
			setup: func() {
				NotFoundPagePath = "testdata/404.html"
			},
			teardown: func() {
				NotFoundPagePath = ""
			},
			want: []byte("Custom 404 page content"),
		},
		{
			name: "Test case 2",
			setup: func() {
				NotFoundPagePath = "non_existing_file"
			},
			teardown: func() {
				NotFoundPagePath = ""
			},
			want: []byte(NotFound),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			tt.setup()
			defer tt.teardown()
			if got := getNotFoundPageContent(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("getNotFoundPageContent() = %v, want %v", got, tt.want)
			}
		})
	}
}
