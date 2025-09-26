package controllers

import (
	"fmt"
	"testing"
)

func TestDoLogin_1(t *testing.T) {
	type args struct {
		username string
		password string
		explicit bool
	}
	tests := []struct {
		name string
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

			self := &LoginController{}
			if got := self.doLogin(tt.args.username, tt.args.password, tt.args.explicit); got != tt.want {
				t.Errorf("doLogin() = %v, want %v", got, tt.want)
			}
		})
	}
}
