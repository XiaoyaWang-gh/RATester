package fake

import (
	"fmt"
	"testing"
)

func TestTraefikV1alpha1_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	cs := &Clientset{}
	traefikv1alpha1 := cs.TraefikV1alpha1()
	if traefikv1alpha1 == nil {
		t.Fatal("TraefikV1alpha1() should not return nil")
	}
}
