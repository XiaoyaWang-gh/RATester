package client

import (
	"fmt"
	"testing"
)

func TestparserRequestBodyFile_1(t *testing.T) {
	type args struct {
		req *Request
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			if err := parserRequestBodyFile(tt.args.req); (err != nil) != tt.wantErr {
				t.Errorf("parserRequestBodyFile() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
