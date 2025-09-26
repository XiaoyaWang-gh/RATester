package hugio

import (
	"fmt"
	"io"
	"strings"
	"testing"
)

func TestReadString_1(t *testing.T) {
	type args struct {
		r io.Reader
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		{
			name: "Test case 1",
			args: args{
				r: strings.NewReader("Hello, World!"),
			},
			want:    "Hello, World!",
			wantErr: false,
		},
		{
			name: "Test case 2",
			args: args{
				r: strings.NewReader(""),
			},
			want:    "",
			wantErr: false,
		},
		{
			name: "Test case 3",
			args: args{
				r: strings.NewReader("Test with error"),
			},
			want:    "Test with error",
			wantErr: false,
		},
		{
			name: "Test case 4",
			args: args{
				r: strings.NewReader(""),
			},
			want:    "",
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			got, err := ReadString(tt.args.r)
			if (err != nil) != tt.wantErr {
				t.Errorf("ReadString() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("ReadString() = %v, want %v", got, tt.want)
			}
		})
	}
}
