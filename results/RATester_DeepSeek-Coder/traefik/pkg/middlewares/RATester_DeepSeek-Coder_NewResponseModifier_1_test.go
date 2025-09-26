package middlewares

import (
	"fmt"
	"net/http"
	"reflect"
	"testing"
)

func TestNewResponseModifier_1(t *testing.T) {
	tests := []struct {
		name     string
		w        http.ResponseWriter
		r        *http.Request
		modifier func(*http.Response) error
		want     *ResponseModifier
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

			got := NewResponseModifier(tt.w, tt.r, tt.modifier)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewResponseModifier() = %v, want %v", got, tt.want)
			}
		})
	}
}
