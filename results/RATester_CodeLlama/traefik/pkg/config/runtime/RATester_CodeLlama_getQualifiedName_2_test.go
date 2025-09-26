package runtime

import (
	"fmt"
	"testing"
)

func TestGetQualifiedName_2(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	var provider string
	var elementName string
	var expected string
	var actual string

	// Test case 1
	provider = "provider1"
	elementName = "element1"
	expected = "element1@provider1"
	actual = getQualifiedName(provider, elementName)
	if actual != expected {
		t.Errorf("getQualifiedName(%s, %s) = %s, expected %s", provider, elementName, actual, expected)
	}

	// Test case 2
	provider = "provider2"
	elementName = "element2"
	expected = "element2@provider2"
	actual = getQualifiedName(provider, elementName)
	if actual != expected {
		t.Errorf("getQualifiedName(%s, %s) = %s, expected %s", provider, elementName, actual, expected)
	}

	// Test case 3
	provider = "provider3"
	elementName = "element3@provider3"
	expected = "element3@provider3"
	actual = getQualifiedName(provider, elementName)
	if actual != expected {
		t.Errorf("getQualifiedName(%s, %s) = %s, expected %s", provider, elementName, actual, expected)
	}

	// Test case 4
	provider = "provider4"
	elementName = "element4@provider4"
	expected = "element4@provider4"
	actual = getQualifiedName(provider, elementName)
	if actual != expected {
		t.Errorf("getQualifiedName(%s, %s) = %s, expected %s", provider, elementName, actual, expected)
	}

	// Test case 5
	provider = "provider5"
	elementName = "element5@provider5"
	expected = "element5@provider5"
	actual = getQualifiedName(provider, elementName)
	if actual != expected {
		t.Errorf("getQualifiedName(%s, %s) = %s, expected %s", provider, elementName, actual, expected)
	}
}
