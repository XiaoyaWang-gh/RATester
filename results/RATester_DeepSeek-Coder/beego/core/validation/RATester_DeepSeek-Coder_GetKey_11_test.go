package validation

import (
	"fmt"
	"regexp"
	"testing"
)

func TestGetKey_11(t *testing.T) {
	type fields struct {
		Match Match
		Key   string
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{
			name: "TestGetKey",
			fields: fields{
				Match: Match{
					Regexp: regexp.MustCompile("^1[3-9]{1}\\d{8}$"),
					Key:    "phone",
				},
				Key: "telephone",
			},
			want: "telephone",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			tel := Tel{
				Match: tt.fields.Match,
				Key:   tt.fields.Key,
			}
			if got := tel.GetKey(); got != tt.want {
				t.Errorf("GetKey() = %v, want %v", got, tt.want)
			}
		})
	}
}
