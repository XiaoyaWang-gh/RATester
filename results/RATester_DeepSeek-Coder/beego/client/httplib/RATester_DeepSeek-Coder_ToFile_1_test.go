package httplib

import (
	"fmt"
	"testing"
)

func TestToFile_1(t *testing.T) {
	type args struct {
		filename string
	}
	tests := []struct {
		name    string
		b       *BeegoHTTPRequest
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

			if err := tt.b.ToFile(tt.args.filename); (err != nil) != tt.wantErr {
				t.Errorf("BeegoHTTPRequest.ToFile() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
