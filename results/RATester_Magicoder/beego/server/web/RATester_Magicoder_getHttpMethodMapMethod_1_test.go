package web

import (
	"fmt"
	"reflect"
	"testing"
)

func TestgetHttpMethodMapMethod_1(t *testing.T) {
	cr := &ControllerRegister{}

	tests := []struct {
		name       string
		httpMethod string
		ctMethod   string
		want       map[string]string
	}{
		{
			name:       "Test with httpMethod != '*' and ctMethod != ''",
			httpMethod: "GET",
			ctMethod:   "POST",
			want:       map[string]string{"GET": "POST"},
		},
		{
			name:       "Test with httpMethod != '*' and ctMethod == ''",
			httpMethod: "GET",
			ctMethod:   "",
			want:       map[string]string{"GET": "GET"},
		},
		{
			name:       "Test with httpMethod == '*' and ctMethod != ''",
			httpMethod: "*",
			ctMethod:   "POST",
			want:       map[string]string{"GET": "POST", "POST": "POST", "PUT": "POST", "DELETE": "POST", "HEAD": "POST", "OPTIONS": "POST", "PATCH": "POST"},
		},
		{
			name:       "Test with httpMethod == '*' and ctMethod == ''",
			httpMethod: "*",
			ctMethod:   "",
			want:       map[string]string{"GET": "GET", "POST": "POST", "PUT": "PUT", "DELETE": "DELETE", "HEAD": "HEAD", "OPTIONS": "OPTIONS", "PATCH": "PATCH"},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			if got := cr.getHttpMethodMapMethod(tt.httpMethod, tt.ctMethod); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("getHttpMethodMapMethod() = %v, want %v", got, tt.want)
			}
		})
	}
}
