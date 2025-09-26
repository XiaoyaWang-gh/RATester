package couchbase

import (
	"fmt"
	"testing"
)

func TestInitOldStyle_1(t *testing.T) {
	type args struct {
		savePath string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "Test with valid savePath",
			args: args{
				savePath: "path1,pool1,bucket1",
			},
			wantErr: false,
		},
		{
			name: "Test with invalid savePath",
			args: args{
				savePath: "",
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			cp := &Provider{}
			if err := cp.initOldStyle(tt.args.savePath); (err != nil) != tt.wantErr {
				t.Errorf("Provider.initOldStyle() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
