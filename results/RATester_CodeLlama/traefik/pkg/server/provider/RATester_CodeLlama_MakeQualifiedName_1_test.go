package provider

import (
	"fmt"
	"testing"
)

func TestMakeQualifiedName_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	var providerName string
	var elementName string
	var expected string
	var actual string

	// CONTEXT_0
	providerName = "providerName"
	elementName = "elementName"
	expected = "elementName@providerName"
	actual = MakeQualifiedName(providerName, elementName)
	if actual != expected {
		t.Errorf("MakeQualifiedName(%s, %s) = %s, expected %s", providerName, elementName, actual, expected)
	}

	// CONTEXT_1
	providerName = "providerName"
	elementName = "elementName"
	expected = "elementName@providerName"
	actual = MakeQualifiedName(providerName, elementName)
	if actual != expected {
		t.Errorf("MakeQualifiedName(%s, %s) = %s, expected %s", providerName, elementName, actual, expected)
	}

	// CONTEXT_2
	providerName = "providerName"
	elementName = "elementName"
	expected = "elementName@providerName"
	actual = MakeQualifiedName(providerName, elementName)
	if actual != expected {
		t.Errorf("MakeQualifiedName(%s, %s) = %s, expected %s", providerName, elementName, actual, expected)
	}

	// CONTEXT_3
	providerName = "providerName"
	elementName = "elementName"
	expected = "elementName@providerName"
	actual = MakeQualifiedName(providerName, elementName)
	if actual != expected {
		t.Errorf("MakeQualifiedName(%s, %s) = %s, expected %s", providerName, elementName, actual, expected)
	}
}
