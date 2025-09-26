package zk

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestInit_12(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	p := &Provider{}
	p.SetDefaults()
	p.Endpoints = []string{"localhost:2181"}
	p.Username = "username"
	p.Password = "password"
	err := p.Init()
	require.NoError(t, err)
}
