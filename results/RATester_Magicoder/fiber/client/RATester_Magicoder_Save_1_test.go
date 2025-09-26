package client

import (
	"fmt"
	"testing"
)

func TestSave_1(t *testing.T) {
	type args struct {
		v any
	}
	tests := []struct {
		name    string
		r       *Response
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

			if err := tt.r.Save(tt.args.v); (err != nil) != tt.wantErr {
				t.Errorf("Response.Save() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
