package resources

import (
	"fmt"
	"testing"
)

func TestcleanKey_3(t *testing.T) {
	type args struct {
		key string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "Test cleanKey with valid key",
			args: args{key: "/test/key"},
			want: "test/key",
		},
		{
			name: "Test cleanKey with empty key",
			args: args{key: ""},
			want: "",
		},
		{
			name: "Test cleanKey with key with leading slash",
			args: args{key: "/test/key/"},
			want: "test/key",
		},
		{
			name: "Test cleanKey with key with trailing slash",
			args: args{key: "test/key/"},
			want: "test/key",
		},
		{
			name: "Test cleanKey with key with multiple slashes",
			args: args{key: "//test//key//"},
			want: "test/key",
		},
		{
			name: "Test cleanKey with key with uppercase letters",
			args: args{key: "TEST/Key"},
			want: "test/key",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			c := &ResourceCache{}
			if got := c.cleanKey(tt.args.key); got != tt.want {
				t.Errorf("cleanKey() = %v, want %v", got, tt.want)
			}
		})
	}
}
