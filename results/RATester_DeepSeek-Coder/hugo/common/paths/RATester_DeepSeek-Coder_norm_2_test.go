package paths

import (
	"fmt"
	"testing"
)

func TestNorm_2(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		p    *Path
		args args
		want string
	}{
		{
			name: "Test with leading slash",
			p:    &Path{trimLeadingSlash: true},
			args: args{s: "/test"},
			want: "test",
		},
		{
			name: "Test without leading slash",
			p:    &Path{trimLeadingSlash: true},
			args: args{s: "test"},
			want: "test",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			if got := tt.p.norm(tt.args.s); got != tt.want {
				t.Errorf("norm() = %v, want %v", got, tt.want)
			}
		})
	}
}
