package crypt

import (
	"fmt"
	"testing"

	"gotest.tools/assert"
)

func TestGetServerName_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	// given
	m := &ClientHelloMsg{
		serverName: "www.example.com",
	}

	// when
	actual := m.GetServerName()

	// then
	assert.Equal(t, "www.example.com", actual)
}
