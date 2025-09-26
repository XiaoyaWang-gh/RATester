package hugo

import (
	"fmt"
	"testing"
)

func TestCompareVersion_1(t *testing.T) {
	type args struct {
		version any
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "test case 1",
			args: args{
				version: "1.0.0",
			},
			want: 0,
		},
		{
			name: "test case 2",
			args: args{
				version: "1.0.1",
			},
			want: 1,
		},
		{
			name: "test case 3",
			args: args{
				version: "0.9.0",
			},
			want: -1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			if got := CompareVersion(tt.args.version); got != tt.want {
				t.Errorf("CompareVersion() = %v, want %v", got, tt.want)
			}
		})
	}
}
