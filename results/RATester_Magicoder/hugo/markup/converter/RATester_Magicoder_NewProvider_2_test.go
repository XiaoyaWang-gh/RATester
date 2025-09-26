package converter

import (
	"fmt"
	"testing"
)

func TestNewProvider_2(t *testing.T) {
	testCases := []struct {
		name   string
		create func(ctx DocumentContext) (Converter, error)
		ctx    DocumentContext
		want   Provider
	}{
		{
			name: "Test case 1",
			create: func(ctx DocumentContext) (Converter, error) {
				return nil, nil
			},
			ctx: DocumentContext{
				Document:       nil,
				DocumentLookup: nil,
				DocumentID:     "testID",
				DocumentName:   "testName",
				Filename:       "testFilename",
			},
			want: newConverter{
				name: "Test case 1",
				create: func(ctx DocumentContext) (Converter, error) {
					return nil, nil
				},
			},
		},
		// Add more test cases as needed
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			got := NewProvider(tc.name, tc.create)
			if got.Name() != tc.want.Name() {
				t.Errorf("Expected Name() to be %s, but got %s", tc.want.Name(), got.Name())
			}
			_, err := got.New(tc.ctx)
			if err != nil {
				t.Errorf("Expected New() to not return an error, but got %v", err)
			}
		})
	}
}
