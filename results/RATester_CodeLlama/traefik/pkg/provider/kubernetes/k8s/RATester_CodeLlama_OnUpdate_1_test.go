package k8s

import (
	"fmt"
	"testing"

	v1 "k8s.io/api/core/v1"
)

func TestOnUpdate_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	// given
	reh := &ResourceEventHandler{}
	oldObj := &v1.Service{}
	newObj := &v1.Service{}

	// when
	reh.OnUpdate(oldObj, newObj)

	// then
	// TODO
}
