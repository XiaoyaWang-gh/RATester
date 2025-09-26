package filecache

import (
	"bytes"
	"errors"
	"fmt"
	"testing"

	"github.com/spf13/afero"
)

func TestGetBytes_1(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		name     string
		id       string
		expected []byte
		err      error
	}{
		{
			name:     "success",
			id:       "test_id",
			expected: []byte("test_data"),
			err:      nil,
		},
		{
			name:     "error",
			id:       "error_id",
			expected: nil,
			err:      errors.New("test error"),
		},
	}

	for _, tc := range testCases {
		tc := tc // capture range variable
		t.Run(tc.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			t.Parallel()

			c := &Cache{
				Fs: afero.NewMemMapFs(),
			}

			afero.WriteFile(c.Fs, tc.id, tc.expected, 0644)

			info, data, err := c.GetBytes(tc.id)

			if !errors.Is(err, tc.err) {
				t.Errorf("expected error %v, got %v", tc.err, err)
			}

			if !bytes.Equal(data, tc.expected) {
				t.Errorf("expected data %v, got %v", tc.expected, data)
			}

			if info.Name != tc.id {
				t.Errorf("expected name %s, got %s", tc.id, info.Name)
			}
		})
	}
}
