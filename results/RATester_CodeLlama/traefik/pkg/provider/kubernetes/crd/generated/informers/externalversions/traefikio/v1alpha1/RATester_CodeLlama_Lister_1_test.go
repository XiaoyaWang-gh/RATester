package v1alpha1

import (
	"fmt"
	"testing"
)

func TestLister_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	f := &traefikServiceInformer{}
	f.Lister()
}
