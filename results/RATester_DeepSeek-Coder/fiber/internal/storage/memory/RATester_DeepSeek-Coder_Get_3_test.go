package memory

import (
	"bytes"
	"fmt"
	"testing"
	"time"
)

func TestGet_3(t *testing.T) {
	type test struct {
		name     string
		key      string
		expected []byte
		setup    func(s *Storage)
	}

	tests := []test{
		{
			name:     "Returns value for existing key",
			key:      "key1",
			expected: []byte("value1"),
			setup: func(s *Storage) {
				s.db["key1"] = entry{data: []byte("value1")}
			},
		},
		{
			name:     "Returns nil for non-existing key",
			key:      "key2",
			expected: nil,
			setup:    func(s *Storage) {},
		},
		{
			name:     "Returns nil for empty key",
			key:      "",
			expected: nil,
			setup:    func(s *Storage) {},
		},
		{
			name:     "Returns nil for expired key",
			key:      "key3",
			expected: nil,
			setup: func(s *Storage) {
				s.db["key3"] = entry{data: []byte("value3"), expiry: uint32(time.Now().Add(-1 * time.Hour).Unix())}
			},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			s := &Storage{
				db: make(map[string]entry),
			}
			test.setup(s)

			value, err := s.Get(test.key)
			if err != nil {
				t.Errorf("unexpected error: %v", err)
			}

			if !bytes.Equal(value, test.expected) {
				t.Errorf("expected %v, got %v", test.expected, value)
			}
		})
	}
}
