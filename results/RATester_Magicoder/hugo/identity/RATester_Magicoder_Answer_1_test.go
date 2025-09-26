package identity

import (
	"fmt"
	"testing"
)

func TestAnswer_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	q := &Question[int]{}

	q.Answer(func() int {
		return 42
	})

	result, ok := q.Result()
	if !ok {
		t.Fatal("Expected question to be answered")
	}

	if result != 42 {
		t.Fatalf("Expected result to be 42, got %v", result)
	}
}
