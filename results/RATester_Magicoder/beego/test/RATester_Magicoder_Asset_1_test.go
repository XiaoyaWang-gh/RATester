package test

import (
	"fmt"
	"testing"
)

func TestAsset_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	tests := []struct {
		name     string
		expected []byte
		err      error
	}{
		{"test.txt", []byte("test content"), nil},
		{"not_found.txt", nil, fmt.Errorf("Asset not_found.txt not found")},
	}

	for _, test := range tests {
		actual, err := Asset(test.name)
		if err != nil && err.Error() != test.err.Error() {
			t.Errorf("Asset(%s) returned error %v, expected %v", test.name, err, test.err)
		}
		if string(actual) != string(test.expected) {
			t.Errorf("Asset(%s) returned %s, expected %s", test.name, actual, test.expected)
		}
	}
}
