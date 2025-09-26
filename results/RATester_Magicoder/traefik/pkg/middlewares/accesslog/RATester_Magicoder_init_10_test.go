package accesslog

import (
	"fmt"
	"testing"
)

func TestInit_10(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	defaultCoreKeys := []string{ServiceAddr, ClientAddr, RequestAddr, GzipRatio, StartLocal, Overhead, RetryAttempts, TLSVersion, TLSCipher, TLSClientSubject}

	for _, k := range defaultCoreKeys {
		if _, ok := allCoreKeys[k]; !ok {
			t.Errorf("Key %s not found in allCoreKeys", k)
		}
	}
}
