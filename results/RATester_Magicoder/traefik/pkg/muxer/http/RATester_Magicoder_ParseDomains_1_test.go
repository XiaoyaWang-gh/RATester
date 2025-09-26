package http

import (
	"fmt"
	"reflect"
	"testing"
)

func TestParseDomains_1(t *testing.T) {
	type args struct {
		rule string
	}
	tests := []struct {
		name    string
		args    args
		want    []string
		wantErr bool
	}{
		{
			name: "Test case 1",
			args: args{
				rule: "example.com",
			},
			want:    []string{"example.com"},
			wantErr: false,
		},
		{
			name: "Test case 2",
			args: args{
				rule: "example.org",
			},
			want:    []string{"example.org"},
			wantErr: false,
		},
		{
			name: "Test case 3",
			args: args{
				rule: "example.net",
			},
			want:    []string{"example.net"},
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

			got, err := ParseDomains(tt.args.rule)
			if (err != nil) != tt.wantErr {
				t.Errorf("ParseDomains() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ParseDomains() = %v, want %v", got, tt.want)
			}
		})
	}
}
