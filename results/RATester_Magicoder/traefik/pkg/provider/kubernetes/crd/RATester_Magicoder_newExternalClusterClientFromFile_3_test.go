package crd

import (
	"fmt"
	"reflect"
	"testing"
)

func TestNewExternalClusterClientFromFile_3(t *testing.T) {
	tests := []struct {
		name    string
		file    string
		want    *clientWrapper
		wantErr bool
	}{
		{
			name: "Test case 1",
			file: "testdata/kubeconfig",
			want: &clientWrapper{
				// Initialize the fields of the clientWrapper struct
			},
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

			got, err := newExternalClusterClientFromFile(tt.file)
			if (err != nil) != tt.wantErr {
				t.Errorf("newExternalClusterClientFromFile() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("newExternalClusterClientFromFile() = %v, want %v", got, tt.want)
			}
		})
	}
}
