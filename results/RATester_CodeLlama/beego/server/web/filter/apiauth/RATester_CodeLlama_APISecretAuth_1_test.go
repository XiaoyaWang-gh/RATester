package apiauth

import (
	"fmt"
	"testing"
)

func TestAPISecretAuth_1(t *testing.T) {
	type args struct {
		f       AppIDToAppSecret
		timeout int
	}
	tests := []struct {
		name string
		args args
	}{
		{"case 1", args{f: func(string) string { return "appsecret" }, timeout: 10}},
		{"case 2", args{f: func(string) string { return "" }, timeout: 10}},
		{"case 3", args{f: func(string) string { return "appsecret" }, timeout: 0}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			APISecretAuth(tt.args.f, tt.args.timeout)
		})
	}
}
