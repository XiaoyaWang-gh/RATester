package version

import (
	"bytes"
	"fmt"
	"io"
	"testing"
)

func TestGetPrint_1(t *testing.T) {
	type args struct {
		wr io.Writer
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "Test case 1",
			args: args{
				wr: &bytes.Buffer{},
			},
			wantErr: false,
		},
		{
			name: "Test case 2",
			args: args{
				wr: &bytes.Buffer{},
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			if err := GetPrint(tt.args.wr); (err != nil) != tt.wantErr {
				t.Errorf("GetPrint() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
