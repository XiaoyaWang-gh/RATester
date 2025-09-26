package mirror

import (
	"fmt"
	"net/http"
	"reflect"
	"testing"
)

func TestHeader_4(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	b := blackHoleResponseWriter{}
	if got := b.Header(); !reflect.DeepEqual(got, http.Header{}) {
		t.Errorf("blackHoleResponseWriter.Header() = %v, want %v", got, http.Header{})
	}
}
