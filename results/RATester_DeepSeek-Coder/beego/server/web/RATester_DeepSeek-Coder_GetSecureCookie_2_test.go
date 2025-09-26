package web

import (
	"fmt"
	"testing"
)

func TestGetSecureCookie_2(t *testing.T) {
	type args struct {
		Secret string
		key    string
	}
	tests := []struct {
		name  string
		c     *Controller
		args  args
		want  string
		want1 bool
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

			got, got1 := tt.c.GetSecureCookie(tt.args.Secret, tt.args.key)
			if got != tt.want {
				t.Errorf("GetSecureCookie() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("GetSecureCookie() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}
