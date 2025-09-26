package identity

import (
	"fmt"
	"testing"
)

func TestNewPredicateIdentity_1(t *testing.T) {
	t.Run("should return a new predicateIdentity with default functions if nil is provided", func(t *testing.T) {

		defer func() {
			if r := recover(); r != nil {
				fmt.Println("Recovered in main", r)
			}
		}()

		predicate := NewPredicateIdentity(nil, nil)
		if predicate.probablyDependent(nil) {
			t.Errorf("expected probablyDependent to be false, got true")
		}
		if predicate.probablyDependency(nil) {
			t.Errorf("expected probablyDependency to be false, got true")
		}
	})

	t.Run("should return a new predicateIdentity with custom functions if provided", func(t *testing.T) {

		defer func() {
			if r := recover(); r != nil {
				fmt.Println("Recovered in main", r)
			}
		}()

		probablyDependent := func(Identity) bool { return true }
		probablyDependency := func(Identity) bool { return true }
		predicate := NewPredicateIdentity(probablyDependent, probablyDependency)
		if !predicate.probablyDependent(nil) {
			t.Errorf("expected probablyDependent to be true, got false")
		}
		if !predicate.probablyDependency(nil) {
			t.Errorf("expected probablyDependency to be true, got false")
		}
	})
}
