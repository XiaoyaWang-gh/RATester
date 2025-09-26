package main

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/traefik/traefik/v3/pkg/config/dynamic"
)

func TestLabelsWrite_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	// given
	element := &dynamic.Configuration{}
	outputDir := "test"

	// when
	err := labelsWrite(outputDir, element)

	// then
	require.NoError(t, err)
}
