package crd

import (
	"fmt"
	"testing"

	"go.uber.org/mock/gomock"
)

func TestGetNodes_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	// arrange
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	// act
	// assert
}
