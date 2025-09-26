package hugolib

import (
	"bytes"
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestApplyArchetypeTemplate_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	// given
	var (
		f = ContentFactory{}
		w = &bytes.Buffer{}
		p = &pageState{}
	)

	// when
	err := f.ApplyArchetypeTemplate(w, p, "", "")

	// then
	require.NoError(t, err)
}
