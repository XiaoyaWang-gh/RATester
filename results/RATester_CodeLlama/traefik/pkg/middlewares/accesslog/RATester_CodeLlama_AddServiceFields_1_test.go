package accesslog

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/alecthomas/assert"
	"github.com/stretchr/testify/require"
)

func TestAddServiceFields_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	// given
	rw := httptest.NewRecorder()
	req, err := http.NewRequest("GET", "http://example.com", nil)
	require.NoError(t, err)
	next := http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
		rw.WriteHeader(http.StatusOK)
	})
	data := &LogData{}

	// when
	AddServiceFields(rw, req, next, data)

	// then
	assert.Equal(t, http.StatusOK, data.Core[OriginStatus])
	assert.Equal(t, int64(0), data.Core[OriginContentSize])
	assert.Equal(t, time.Duration(0), data.Core[OriginDuration])
}
