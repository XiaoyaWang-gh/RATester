package allconfig

import (
	"fmt"
	"testing"
)

func TestIsUglyURLs_1(t *testing.T) {
	type args struct {
		section string
	}
	tests := []struct {
		name string
		c    ConfigLanguage
		args args
		want bool
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

			if got := tt.c.IsUglyURLs(tt.args.section); got != tt.want {
				t.Errorf("IsUglyURLs() = %v, want %v", got, tt.want)
			}
		})
	}
}
