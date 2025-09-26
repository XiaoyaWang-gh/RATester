package alils

import (
	"fmt"
	"testing"
)

func TestNewAliLS_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	alils := NewAliLS()
	if alils == nil {
		t.Errorf("NewAliLS() = %v, want %v", alils, nil)
	}
}
