package utils

import (
	"fmt"
	"io"
	"testing"
)

func TestQuotePrintEncode_1(t *testing.T) {
	type args struct {
		w io.Writer
		s string
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

			if err := quotePrintEncode(tt.args.w, tt.args.s); (err != nil) != tt.wantErr {
				t.Errorf("quotePrintEncode() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
