package keyauth

import (
	"fmt"
	"reflect"
	"testing"
)

func TestKeyFromHeader_1(t *testing.T) {
	type args struct {
		header     string
		authScheme string
	}
	tests := []struct {
		name string
		args args
		want KeyLookupFunc
	}{
		{
			name: "Test Case 1",
			args: args{
				header:     "Authorization",
				authScheme: "Bearer",
			},
			want: KeyFromHeader("Authorization", "Bearer"),
		},
		{
			name: "Test Case 2",
			args: args{
				header:     "API-Key",
				authScheme: "",
			},
			want: KeyFromHeader("API-Key", ""),
		},
		{
			name: "Test Case 3",
			args: args{
				header:     "X-API-Key",
				authScheme: "Custom",
			},
			want: KeyFromHeader("X-API-Key", "Custom"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			if got := KeyFromHeader(tt.args.header, tt.args.authScheme); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("KeyFromHeader() = %v, want %v", got, tt.want)
			}
		})
	}
}
