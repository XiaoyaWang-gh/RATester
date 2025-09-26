package httplib

import (
	"fmt"
	"testing"
)

func TestToValue_1(t *testing.T) {
	type args struct {
		value interface{}
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

			if err := tt.b.ToValue(tt.args.value); (err != nil) != tt.wantErr {
				t.Errorf("BeegoHTTPRequest.ToValue() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
