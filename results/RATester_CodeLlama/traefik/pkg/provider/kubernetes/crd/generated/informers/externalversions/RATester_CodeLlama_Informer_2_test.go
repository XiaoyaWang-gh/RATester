package externalversions

import (
	"fmt"
	"testing"
)

func TestInformer_2(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	f := &genericInformer{}
	f.Informer()
}
