package test

import (
	"fmt"
	"reflect"
	"testing"
)

func TestViewsHeaderTplBytes_1(t *testing.T) {
	tests := []struct {
		name    string
		want    []byte
		wantErr bool
	}{
		{
			name:    "Test case 1",
			want:    []byte("<html><head><title>{{.Title}}</title></head><body>{{.Body}}</body></html>"),
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

			got, err := viewsHeaderTplBytes()
			if (err != nil) != tt.wantErr {
				t.Errorf("viewsHeaderTplBytes() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("viewsHeaderTplBytes() = %v, want %v", got, tt.want)
			}
		})
	}
}
