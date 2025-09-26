package client

import (
	"fmt"
	"testing"
)

func TestReleaseCookieJar_1(t *testing.T) {
	t.Parallel()

	type args struct {
		c *CookieJar
	}

	tests := []struct {
		name string
		args args
	}{
		{
			name: "TestReleaseCookieJar",
			args: args{
				c: &CookieJar{},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			ReleaseCookieJar(tt.args.c)
		})
	}
}
