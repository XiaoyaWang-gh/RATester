package template

import (
	"bytes"
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestFuncs_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	// given
	var tpl *Template
	var err error
	var funcMap FuncMap
	// when
	funcMap = FuncMap{"foo": func() string { return "bar" }}
	tpl, err = New("foo").Funcs(funcMap).Parse("{{foo}}")
	// then
	require.NoError(t, err)
	require.NotNil(t, tpl)
	var buf bytes.Buffer
	err = tpl.Execute(&buf, nil)
	require.NoError(t, err)
	require.Equal(t, "bar", buf.String())
}
