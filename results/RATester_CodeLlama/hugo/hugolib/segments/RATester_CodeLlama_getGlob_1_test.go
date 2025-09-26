package segments

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/gobwas/glob"
)

func TestGetGlob_1(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name    string
		args    args
		want    glob.Glob
		wantErr bool
	}{
		{
			name: "empty string",
			args: args{
				s: "",
			},
			want:    nil,
			wantErr: false,
		},
		{
			name: "valid glob",
			args: args{
				s: "*.txt",
			},
			want:    nil,
			wantErr: false,
		},
		{
			name: "invalid glob",
			args: args{
				s: "*.txt*",
			},
			want:    nil,
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

			got, err := getGlob(tt.args.s)
			if (err != nil) != tt.wantErr {
				t.Errorf("getGlob() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("getGlob() = %v, want %v", got, tt.want)
			}
		})
	}
}
