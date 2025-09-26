package install

import (
	"fmt"
	"testing"
)

func TestDownloadLatest_1(t *testing.T) {
	type args struct {
		bin string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "Test downloadLatest function",
			args: args{bin: "server"},
			want: "expected result",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			if got := downloadLatest(tt.args.bin); got != tt.want {
				t.Errorf("downloadLatest() = %v, want %v", got, tt.want)
			}
		})
	}
}
