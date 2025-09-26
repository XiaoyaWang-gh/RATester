package fake

import (
	"context"
	"fmt"
	"reflect"
	"testing"

	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func TestDelete_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	// arrange
	var ctx context.Context
	var name string
	var opts v1.DeleteOptions
	var expectedError error
	var fakeTraefikServices *FakeTraefikServices
	var actualError error

	// act
	actualError = fakeTraefikServices.Delete(ctx, name, opts)

	// assert
	if !reflect.DeepEqual(actualError, expectedError) {
		t.Errorf("actualError=%v, expectedError=%v", actualError, expectedError)
	}
}
