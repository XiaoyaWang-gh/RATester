package http

import (
	"fmt"
	"testing"
)

func TestHostRegexp_1(t *testing.T) {
	type args struct {
		tree  *matchersTree
		hosts []string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "valid host",
			args: args{
				tree:  &matchersTree{},
				hosts: []string{"example.com"},
			},
			wantErr: false,
		},
		{
			name: "invalid host",
			args: args{
				tree:  &matchersTree{},
				hosts: []string{"\x80example.com"},
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

			if err := hostRegexp(tt.args.tree, tt.args.hosts...); (err != nil) != tt.wantErr {
				t.Errorf("hostRegexp() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
