package session

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestSave_2(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	// given
	s := &Session{
		data: &data{
			Data: map[string]any{
				"key": "value",
			},
		},
	}

	// when
	err := s.Save()

	// then
	require.NoError(t, err)
}
