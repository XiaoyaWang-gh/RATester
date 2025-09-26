package context

import (
	"fmt"
	"testing"
)

func TestCookie_2(t *testing.T) {
	type args struct {
		name   string
		value  string
		others []interface{}
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "TestCookie_0",
			args: args{
				name:  "test_cookie",
				value: "test_value",
				others: []interface{}{
					3600,
					"/test_path",
					"",
					true,
					true,
					"Lax",
				},
			},
			want: "test_cookie=test_value; Expires=Mon, 02 Jan 2006 15:04:05 MST; Max-Age=3600; Path=/test_path; Secure; HttpOnly; SameSite=Lax",
		},
		// Add more test cases here
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			output := &BeegoOutput{
				Context: &Context{
					ResponseWriter: &Response{},
				},
			}
			output.Cookie(tt.args.name, tt.args.value, tt.args.others...)
			if got := output.Context.ResponseWriter.Header().Get("Set-Cookie"); got != tt.want {
				t.Errorf("Cookie() = %v, want %v", got, tt.want)
			}
		})
	}
}
