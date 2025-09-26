package mirror

import (
	"fmt"
	"net/http"
	"reflect"
	"testing"
)

func TestGetActiveMirrors_1(t *testing.T) {
	tests := []struct {
		name string
		m    *Mirroring
		want []http.Handler
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			if got := tt.m.getActiveMirrors(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Mirroring.getActiveMirrors() = %v, want %v", got, tt.want)
			}
		})
	}
}
