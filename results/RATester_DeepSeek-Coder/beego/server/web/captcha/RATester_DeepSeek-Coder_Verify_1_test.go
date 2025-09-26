package captcha

import (
	"fmt"
	"testing"
)

func TestVerify_1(t *testing.T) {
	type args struct {
		id        string
		challenge string
	}
	tests := []struct {
		name        string
		c           *Captcha
		args        args
		wantSuccess bool
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

			if gotSuccess := tt.c.Verify(tt.args.id, tt.args.challenge); gotSuccess != tt.wantSuccess {
				t.Errorf("Verify() = %v, want %v", gotSuccess, tt.wantSuccess)
			}
		})
	}
}
