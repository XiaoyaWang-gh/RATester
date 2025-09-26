package media

import (
	"fmt"
	"reflect"
	"testing"
)

func TestGetFirstBySuffix_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	// Arrange
	var tt Types
	var suffix string
	var expected Type
	var expectedSuffixInfo SuffixInfo
	var expectedFound bool
	var actual Type
	var actualSuffixInfo SuffixInfo
	var actualFound bool

	// Act
	// Assert
	// Arrange
	tt = append(tt, Type{
		Type: "application/rss+xml",
	})
	suffix = "xml"
	expected = Type{
		Type: "application/rss+xml",
	}
	expectedSuffixInfo = SuffixInfo{
		Suffix: "xml",
	}
	expectedFound = true
	actual, actualSuffixInfo, actualFound = tt.GetFirstBySuffix(suffix)
	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("expected %v, but got %v", expected, actual)
	}
	if !reflect.DeepEqual(expectedSuffixInfo, actualSuffixInfo) {
		t.Errorf("expected %v, but got %v", expectedSuffixInfo, actualSuffixInfo)
	}
	if expectedFound != actualFound {
		t.Errorf("expected %v, but got %v", expectedFound, actualFound)
	}
}
