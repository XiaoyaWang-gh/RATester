package validation

import (
	"fmt"
	"regexp"
	"testing"
)

func TestGetKey_4(t *testing.T) {
	type fields struct {
		Regexp *regexp.Regexp
		Key    string
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{
			name: "TestGetKey",
			fields: fields{
				Regexp: regexp.MustCompile("^[a-zA-Z0-9]*$"),
				Key:    "testKey",
			},
			want: "testKey",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			m := Match{
				Regexp: tt.fields.Regexp,
				Key:    tt.fields.Key,
			}
			if got := m.GetKey(); got != tt.want {
				t.Errorf("GetKey() = %v, want %v", got, tt.want)
			}
		})
	}
}
