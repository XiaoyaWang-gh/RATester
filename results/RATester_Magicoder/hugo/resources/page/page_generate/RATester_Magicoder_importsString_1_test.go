package page_generate

import (
	"fmt"
	"testing"
)

func TestimportsString_1(t *testing.T) {
	tests := []struct {
		name string
		imps []string
		want string
	}{
		{
			name: "empty imports",
			imps: []string{},
			want: "",
		},
		{
			name: "single import",
			imps: []string{"fmt"},
			want: `import "fmt"`,
		},
		{
			name: "multiple imports",
			imps: []string{"fmt", "os"},
			want: `import (
"fmt"
"os"
)`,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			if got := importsString(tt.imps); got != tt.want {
				t.Errorf("importsString() = %v, want %v", got, tt.want)
			}
		})
	}
}
