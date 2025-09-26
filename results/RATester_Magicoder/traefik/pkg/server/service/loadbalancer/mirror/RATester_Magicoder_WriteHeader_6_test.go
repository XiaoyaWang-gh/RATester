package mirror

import (
	"fmt"
	"net/http"
	"testing"
)

func TestWriteHeader_6(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	b := blackHoleResponseWriter{}
	b.WriteHeader(http.StatusOK)
}
