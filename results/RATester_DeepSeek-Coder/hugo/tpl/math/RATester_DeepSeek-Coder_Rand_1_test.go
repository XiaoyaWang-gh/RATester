package math

import (
	"fmt"
	"math/rand"
	"testing"
	"time"
)

func TestRand_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	ns := &Namespace{}

	rand.Seed(time.Now().UnixNano())
	randValue := ns.Rand()

	if randValue < 0 || randValue > 1 {
		t.Errorf("Expected rand value to be between 0 and 1, got %f", randValue)
	}
}
