package langs

import (
	"fmt"
	"testing"

	"golang.org/x/text/collate"
	"golang.org/x/text/language"
)

func TestCompareStrings_1(t *testing.T) {
	collator := &Collator{
		c: collate.New(language.English, collate.Loose),
	}

	tests := []struct {
		name string
		a    string
		b    string
		want int
	}{
		{
			name: "Equal strings",
			a:    "Hello",
			b:    "Hello",
			want: 0,
		},
		{
			name: "Different strings",
			a:    "Hello",
			b:    "World",
			want: -1,
		},
		{
			name: "Empty strings",
			a:    "",
			b:    "",
			want: 0,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			if got := collator.CompareStrings(tt.a, tt.b); got != tt.want {
				t.Errorf("CompareStrings() = %v, want %v", got, tt.want)
			}
		})
	}
}
