package connectionheader

import (
	"fmt"
	"net/http"
	"reflect"
	"testing"
)

func TestRemove_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	req := &http.Request{
		Header: http.Header{
			"Connection": []string{"keep-alive"},
			"Upgrade":    []string{"websocket"},
		},
	}
	want := &http.Request{
		Header: http.Header{
			"Upgrade": []string{"websocket"},
		},
	}
	got := Remove(req)
	if !reflect.DeepEqual(got, want) {
		t.Errorf("Remove() = %v, want %v", got, want)
	}
}
