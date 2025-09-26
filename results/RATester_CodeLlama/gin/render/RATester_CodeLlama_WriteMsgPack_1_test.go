package render

import (
	"fmt"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestWriteMsgPack_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	// given
	w := httptest.NewRecorder()
	obj := struct {
		Name string
	}{
		Name: "test",
	}

	// when
	err := WriteMsgPack(w, obj)

	// then
	require.NoError(t, err)
	require.Equal(t, msgpackContentType, w.Header().Get("Content-Type"))
	require.Equal(t, "test", w.Body.String())
}
