package yaml

import (
	"errors"
	"fmt"
	"testing"
)

func TestBool_1(t *testing.T) {
	t.Parallel()

	cc := &ConfigContainer{
		data: map[string]interface{}{
			"key1": "true",
			"key2": "false",
			"key3": "invalid",
		},
	}

	tests := []struct {
		name     string
		key      string
		expected bool
		err      error
	}{
		{"valid true", "key1", true, nil},
		{"valid false", "key2", false, nil},
		{"invalid", "key3", false, errors.New("invalid value")},
		{"not exist", "key4", false, errors.New("key not exist")},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			got, err := cc.Bool(tt.key)
			if err != nil && err.Error() != tt.err.Error() {
				t.Errorf("expected error %v, got %v", tt.err, err)
			}
			if got != tt.expected {
				t.Errorf("expected %v, got %v", tt.expected, got)
			}
		})
	}
}
