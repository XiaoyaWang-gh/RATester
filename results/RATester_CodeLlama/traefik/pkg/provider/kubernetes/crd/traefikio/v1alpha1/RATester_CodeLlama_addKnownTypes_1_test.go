package v1alpha1

import (
	"fmt"
	"testing"

	"k8s.io/apimachinery/pkg/runtime"
)

func TestAddKnownTypes_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	scheme := runtime.NewScheme()
	err := addKnownTypes(scheme)
	if err != nil {
		t.Errorf("addKnownTypes() error = %v", err)
		return
	}
	if scheme.PrioritizedVersionsAllGroups() == nil {
		t.Errorf("addKnownTypes() error = %v", err)
		return
	}
}
