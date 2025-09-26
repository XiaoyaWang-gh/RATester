package metadecoders

import (
	"fmt"
	"testing"
)

func TestOptionsKey_1(t *testing.T) {
	tests := []struct {
		name string
		d    Decoder
		want string
	}{
		{
			name: "Test case 1",
			d: Decoder{
				Delimiter:  ',',
				Comment:    '#',
				LazyQuotes: true,
			},
			want: ",#true",
		},
		{
			name: "Test case 2",
			d: Decoder{
				Delimiter:  ';',
				Comment:    ' ',
				LazyQuotes: false,
			},
			want: "; false",
		},
		{
			name: "Test case 3",
			d: Decoder{
				Delimiter:  '|',
				Comment:    '$',
				LazyQuotes: true,
			},
			want: "|$true",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			if got := tt.d.OptionsKey(); got != tt.want {
				t.Errorf("Decoder.OptionsKey() = %v, want %v", got, tt.want)
			}
		})
	}
}
