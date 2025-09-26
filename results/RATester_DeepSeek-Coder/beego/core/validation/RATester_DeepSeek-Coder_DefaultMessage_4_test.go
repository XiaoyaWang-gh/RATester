package validation

import (
	"fmt"
	"regexp"
	"testing"
)

func TestDefaultMessage_4(t *testing.T) {
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
			name: "Test DefaultMessage",
			fields: fields{
				Match: Match{
					Regexp: regexp.MustCompile(`^1[3-9]\d{9}$`),
					Key:    "Tel",
				},
				Key: "Tel",
			},
			want: "Tel",
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
			if got := tel.DefaultMessage(); got != tt.want {
				t.Errorf("DefaultMessage() = %v, want %v", got, tt.want)
			}
		})
	}
}
