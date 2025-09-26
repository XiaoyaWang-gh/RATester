package pageparser

import (
	"fmt"
	"io"
	"reflect"
	"strings"
	"testing"
)

func TestparseSection_1(t *testing.T) {
	type args struct {
		r     io.Reader
		cfg   Config
		start stateFunc
	}
	tests := []struct {
		name    string
		args    args
		want    Result
		wantErr bool
	}{
		{
			name: "Test Case 1",
			args: args{
				r:   strings.NewReader("test content"),
				cfg: Config{NoFrontMatter: true},
				start: func(l *pageLexer) stateFunc {
					return nil
				},
			},
			want:    nil,
			wantErr: false,
		},
		{
			name: "Test Case 2",
			args: args{
				r:   strings.NewReader("test content"),
				cfg: Config{NoFrontMatter: false},
				start: func(l *pageLexer) stateFunc {
					return nil
				},
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

			got, err := parseSection(tt.args.r, tt.args.cfg, tt.args.start)
			if (err != nil) != tt.wantErr {
				t.Errorf("parseSection() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("parseSection() = %v, want %v", got, tt.want)
			}
		})
	}
}
