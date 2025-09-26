package commands

import (
	"fmt"
	"reflect"
	"regexp"
	"testing"
)

func TestfilterIrrelevant_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	detector := &fileChangeDetector{
		irrelevantRe: regexp.MustCompile(`^\.`),
	}

	input := []string{".hidden", "notHidden"}
	expected := []string{"notHidden"}

	result := detector.filterIrrelevant(input)

	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Expected %v, got %v", expected, result)
	}
}
