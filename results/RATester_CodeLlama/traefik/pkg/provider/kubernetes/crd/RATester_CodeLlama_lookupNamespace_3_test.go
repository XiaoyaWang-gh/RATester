package crd

import (
	"fmt"
	"testing"

	"github.com/alecthomas/assert"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func TestLookupNamespace_3(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	c := &clientWrapper{}
	c.isNamespaceAll = true
	ns := "test"
	assert.Equal(t, metav1.NamespaceAll, c.lookupNamespace(ns))
}
