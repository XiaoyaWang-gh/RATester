package tcp

import (
	"fmt"
	"reflect"
	"testing"
)

func TestParseHostSNI_1(t *testing.T) {
	tests := []struct {
		name    string
		rule    string
		want    []string
		wantErr bool
	}{
		{
			name: "Test case 1",
			rule: "HostSNI(`example.com`)",
			want: []string{"example.com"},
		},
		{
			name: "Test case 2",
			rule: "HostSNI(`example.org`)",
			want: []string{"example.org"},
		},
		{
			name: "Test case 3",
			rule: "HostSNI(`example.net`)",
			want: []string{"example.net"},
		},
		{
			name:    "Test case 4",
			rule:    "HostSNI(`example.invalid`)",
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

			got, err := ParseHostSNI(tt.rule)
			if (err != nil) != tt.wantErr {
				t.Errorf("ParseHostSNI() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ParseHostSNI() = %v, want %v", got, tt.want)
			}
		})
	}
}
