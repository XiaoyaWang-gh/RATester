package paths

import (
	"fmt"
	"testing"
)

func TestAbsPathify_1(t *testing.T) {
	type args struct {
		inPath string
	}
	tests := []struct {
		name string
		p    *Paths
		args args
		want string
	}{
		{
			name: "Test case 1",
			p:    &Paths{},
			args: args{
				inPath: "test/path",
			},
			want: "absolute/path/test/path",
		},
		{
			name: "Test case 2",
			p:    &Paths{},
			args: args{
				inPath: "another/test/path",
			},
			want: "absolute/path/another/test/path",
		},
		// Add more test cases as needed
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			if got := tt.p.AbsPathify(tt.args.inPath); got != tt.want {
				t.Errorf("AbsPathify() = %v, want %v", got, tt.want)
			}
		})
	}
}
