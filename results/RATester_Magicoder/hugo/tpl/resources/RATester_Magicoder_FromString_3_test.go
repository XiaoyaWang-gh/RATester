package resources

import (
	"errors"
	"fmt"
	"testing"

	"github.com/gohugoio/hugo/resources/resource_factories/create"
)

func TestFromString_3(t *testing.T) {
	ns := &Namespace{
		createClient: &create.Client{},
	}

	tests := []struct {
		name        string
		targetPath  any
		content     any
		expectedErr error
	}{
		{
			name:       "valid input",
			targetPath: "valid/path",
			content:    "valid content",
		},
		{
			name:        "invalid target path",
			targetPath:  123,
			content:     "valid content",
			expectedErr: errors.New("unable to cast 123 of type int to string"),
		},
		{
			name:        "invalid content",
			targetPath:  "valid/path",
			content:     123,
			expectedErr: errors.New("unable to cast 123 of type int to string"),
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			_, err := ns.FromString(test.targetPath, test.content)
			if err != nil {
				if test.expectedErr == nil {
					t.Errorf("unexpected error: %v", err)
				} else if err.Error() != test.expectedErr.Error() {
					t.Errorf("expected error: %v, got: %v", test.expectedErr, err)
				}
			}
		})
	}
}
