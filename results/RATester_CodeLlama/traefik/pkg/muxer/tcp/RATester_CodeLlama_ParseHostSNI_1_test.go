package tcp

import (
	"fmt"
	"reflect"
	"testing"
)

func TestParseHostSNI_1(t *testing.T) {
	var tests = []struct {
		rule string
		want []string
	}{
		{
			rule: "HostSNI(example.com)",
			want: []string{"example.com"},
		},
		{
			rule: "HostSNI(example.com,example.org)",
			want: []string{"example.com", "example.org"},
		},
		{
			rule: "HostSNI(example.com,example.org,example.net)",
			want: []string{"example.com", "example.org", "example.net"},
		},
	}

	for _, tt := range tests {
		t.Run(tt.rule, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			got, err := ParseHostSNI(tt.rule)
			if err != nil {
				t.Fatal(err)
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ParseHostSNI() = %v, want %v", got, tt.want)
			}
		})
	}
}
