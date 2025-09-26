package test

import (
	"fmt"
	"reflect"
	"testing"
)

func TestViewsIndexTplBytes_1(t *testing.T) {
	tests := []struct {
		name    string
		want    []byte
		wantErr bool
	}{
		{
			name:    "Test case 1",
			want:    []byte("<html><body><h1>Hello, world</h1></body></html>"),
			wantErr: false,
		},
		// Add more test cases as needed
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			got, err := viewsIndexTplBytes()
			if (err != nil) != tt.wantErr {
				t.Errorf("viewsIndexTplBytes() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("viewsIndexTplBytes() = %v, want %v", got, tt.want)
			}
		})
	}
}
