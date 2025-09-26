package integrity

import (
	"fmt"
	"testing"
)

func TestIntegrity_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	var algo string
	var sum []byte
	var expected string
	var actual string
	// Test case 1:
	algo = "sha256"
	sum = []byte("1234567890")
	expected = "sha256-47DEQpj8HBSa-_TImW-5JCeuQeRkm5NMpJWZG3hSuFU="
	actual = integrity(algo, sum)
	if actual != expected {
		t.Errorf("Testcase 1 failed. Got %v, expected %v", actual, expected)
	}
	// Test case 2:
	algo = "sha256"
	sum = []byte("1234567890")
	expected = "sha256-47DEQpj8HBSa-_TImW-5JCeuQeRkm5NMpJWZG3hSuFU="
	actual = integrity(algo, sum)
	if actual != expected {
		t.Errorf("Testcase 2 failed. Got %v, expected %v", actual, expected)
	}
	// Test case 3:
	algo = "sha256"
	sum = []byte("1234567890")
	expected = "sha256-47DEQpj8HBSa-_TImW-5JCeuQeRkm5NMpJWZG3hSuFU="
	actual = integrity(algo, sum)
	if actual != expected {
		t.Errorf("Testcase 3 failed. Got %v, expected %v", actual, expected)
	}
}
