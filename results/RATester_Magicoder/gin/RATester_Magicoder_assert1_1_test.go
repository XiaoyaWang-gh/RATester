package gin

import (
	"fmt"
	"testing"
)

func Testassert1_1(t *testing.T) {
	tests := []struct {
		name   string
		guard  bool
		text   string
		expect string
	}{
		{
			name:   "Test case 1",
			guard:  true,
			text:   "This is a test",
			expect: "",
		},
		{
			name:   "Test case 2",
			guard:  false,
			text:   "This is a test",
			expect: "This is a test",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			defer func() {
				if r := recover(); r != nil {
					if r.(string) != tt.expect {
						t.Errorf("assert1() = %v, want %v", r, tt.expect)
					}
				}
			}()
			assert1(tt.guard, tt.text)
		})
	}
}
