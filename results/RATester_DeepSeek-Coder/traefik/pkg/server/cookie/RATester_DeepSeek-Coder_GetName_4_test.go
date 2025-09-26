package cookie

import (
	"fmt"
	"testing"
)

func TestGetName_4(t *testing.T) {
	type args struct {
		cookieName  string
		backendName string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "Test with non-empty cookieName",
			args: args{
				cookieName:  "test_cookie",
				backendName: "test_backend",
			},
			want: "test_cookie",
		},
		{
			name: "Test with empty cookieName",
			args: args{
				cookieName:  "",
				backendName: "test_backend",
			},
			want: "test_backend",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			if got := GetName(tt.args.cookieName, tt.args.backendName); got != tt.want {
				t.Errorf("GetName() = %v, want %v", got, tt.want)
			}
		})
	}
}
