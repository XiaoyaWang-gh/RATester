package config

import (
	"fmt"
	"testing"
)

func TestGetdata_3(t *testing.T) {
	type args struct {
		key string
	}
	tests := []struct {
		name string
		c    *IniConfigContainer
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

			if got := tt.c.getdata(tt.args.key); got != tt.want {
				t.Errorf("IniConfigContainer.getdata() = %v, want %v", got, tt.want)
			}
		})
	}
}
