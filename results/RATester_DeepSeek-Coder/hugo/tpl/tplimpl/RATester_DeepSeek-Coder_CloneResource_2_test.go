package tplimpl

import (
	"errors"
	"fmt"
	"testing"

	"github.com/gohugoio/hugo/deps"
)

func TestCloneResource_2(t *testing.T) {
	testCases := []struct {
		name     string
		dst      *deps.Deps
		src      *deps.Deps
		expected error
	}{
		{
			name:     "success",
			dst:      &deps.Deps{},
			src:      &deps.Deps{},
			expected: nil,
		},
		{
			name:     "failure",
			dst:      &deps.Deps{},
			src:      &deps.Deps{},
			expected: errors.New("some error"),
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			tp := &TemplateProvider{}
			err := tp.CloneResource(tc.dst, tc.src)
			if err != tc.expected {
				t.Errorf("Expected error %v, got %v", tc.expected, err)
			}
		})
	}
}
