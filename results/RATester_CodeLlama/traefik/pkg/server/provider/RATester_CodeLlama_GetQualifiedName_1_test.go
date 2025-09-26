package provider

import (
	"context"
	"fmt"
	"testing"
)

func TestGetQualifiedName_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	ctx := context.Background()
	elementName := "test"
	qualifiedName := GetQualifiedName(ctx, elementName)
	if qualifiedName != elementName {
		t.Errorf("GetQualifiedName() = %v, want %v", qualifiedName, elementName)
	}
}
