package captcha

import (
	"fmt"
	"html/template"
	"testing"
)

func TestCreateCaptchaHTML_1(t *testing.T) {
	type args struct {
		c *Captcha
	}
	tests := []struct {
		name string
		args args
		want template.HTML
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

			if got := tt.args.c.CreateCaptchaHTML(); got != tt.want {
				t.Errorf("Captcha.CreateCaptchaHTML() = %v, want %v", got, tt.want)
			}
		})
	}
}
