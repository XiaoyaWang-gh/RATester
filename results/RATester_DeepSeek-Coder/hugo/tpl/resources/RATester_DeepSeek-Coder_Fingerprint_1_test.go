package resources

import (
	"errors"
	"fmt"
	"testing"

	"github.com/gohugoio/hugo/resources/resource"
	"github.com/gohugoio/hugo/resources/resource_transformers/integrity"
)

func TestFingerprint_1(t *testing.T) {
	t.Parallel()

	type testCase struct {
		name        string
		args        []any
		expectedErr error
	}

	testCases := []testCase{
		{
			name:        "no args",
			args:        []any{},
			expectedErr: errors.New("must provide a Resource object"),
		},
		{
			name: "too many args",
			args: []any{
				new(resource.Resource),
				"sha256",
				"sha512",
			},
			expectedErr: errors.New("must not provide more arguments than Resource and hash algorithm"),
		},
		{
			name: "invalid hash algorithm",
			args: []any{
				new(resource.Resource),
				123,
			},
			expectedErr: errors.New("can't cast 123 of type int to string"),
		},
		{
			name: "invalid resource",
			args: []any{
				"not a resource",
			},
			expectedErr: fmt.Errorf("%T can not be transformed", "not a resource"),
		},
		{
			name: "valid case",
			args: []any{
				new(resource.Resource),
			},
			expectedErr: nil,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			ns := &Namespace{
				integrityClient: &integrity.Client{},
			}

			_, err := ns.Fingerprint(tc.args...)
			if err != nil && tc.expectedErr != nil && err.Error() != tc.expectedErr.Error() {
				t.Errorf("expected error %v, got %v", tc.expectedErr, err)
			}

			if err != tc.expectedErr {
				t.Errorf("expected error %v, got %v", tc.expectedErr, err)
			}
		})
	}
}
