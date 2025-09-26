package cors

import (
	"fmt"
	"testing"
)

func TestMatch_3(t *testing.T) {
	type args struct {
		o string
	}
	tests := []struct {
		name string
		s    subdomain
		args args
		want bool
	}{
		{
			name: "Test case 1",
			s:    subdomain{prefix: "www", suffix: ".example.com"},
			args: args{o: "www.example.com"},
			want: true,
		},
		{
			name: "Test case 2",
			s:    subdomain{prefix: "www", suffix: ".example.com"},
			args: args{o: "www.example.com."},
			want: true,
		},
		{
			name: "Test case 3",
			s:    subdomain{prefix: "www", suffix: ".example.com"},
			args: args{o: "www.example.com.au"},
			want: false,
		},
		{
			name: "Test case 4",
			s:    subdomain{prefix: "www", suffix: ".example.com"},
			args: args{o: "example.com"},
			want: false,
		},
		{
			name: "Test case 5",
			s:    subdomain{prefix: "www", suffix: ".example.com"},
			args: args{o: "www.example.com.www"},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			if got := tt.s.match(tt.args.o); got != tt.want {
				t.Errorf("subdomain.match() = %v, want %v", got, tt.want)
			}
		})
	}
}
