package kinds

import (
	"fmt"
	"testing"
)

func TestIsDeprecatedAndReplacedWith_1(t *testing.T) {
	tests := []struct {
		name string
		arg  string
		want string
	}{
		{
			name: "Test case 1",
			arg:  "TaxonomyTerm",
			want: KindTaxonomy,
		},
		{
			name: "Test case 2",
			arg:  "taxonomyterm",
			want: KindTaxonomy,
		},
		{
			name: "Test case 3",
			arg:  "Other",
			want: "",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			if got := IsDeprecatedAndReplacedWith(tt.arg); got != tt.want {
				t.Errorf("IsDeprecatedAndReplacedWith() = %v, want %v", got, tt.want)
			}
		})
	}
}
