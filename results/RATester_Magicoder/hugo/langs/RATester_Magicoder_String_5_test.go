package langs

import (
	"fmt"
	"testing"
)

func TestString_5(t *testing.T) {
	type fields struct {
		Lang string
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{
			name: "Test Case 1",
			fields: fields{
				Lang: "en",
			},
			want: "en",
		},
		{
			name: "Test Case 2",
			fields: fields{
				Lang: "no",
			},
			want: "no",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			l := &Language{
				Lang: tt.fields.Lang,
			}
			if got := l.String(); got != tt.want {
				t.Errorf("Language.String() = %v, want %v", got, tt.want)
			}
		})
	}
}
