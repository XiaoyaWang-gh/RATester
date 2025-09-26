package captcha

import (
	"fmt"
	"testing"
)

func TestKey_2(t *testing.T) {
	type args struct {
		id string
	}
	tests := []struct {
		name string
		c    *Captcha
		args args
		want string
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

			if got := tt.c.key(tt.args.id); got != tt.want {
				t.Errorf("key() = %v, want %v", got, tt.want)
			}
		})
	}
}
