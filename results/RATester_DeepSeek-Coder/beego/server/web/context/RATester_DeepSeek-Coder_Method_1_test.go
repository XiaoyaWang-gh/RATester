package context

import (
	"fmt"
	"net/http"
	"testing"
)

func TestMethod_1(t *testing.T) {
	type testStruct struct {
		input *BeegoInput
		want  string
	}

	tests := []testStruct{
		{&BeegoInput{Context: &Context{Request: &http.Request{Method: "GET"}}}, "GET"},
		{&BeegoInput{Context: &Context{Request: &http.Request{Method: "POST"}}}, "POST"},
		{&BeegoInput{Context: &Context{Request: &http.Request{Method: "PUT"}}}, "PUT"},
		{&BeegoInput{Context: &Context{Request: &http.Request{Method: "DELETE"}}}, "DELETE"},
		{&BeegoInput{Context: &Context{Request: &http.Request{Method: "PATCH"}}}, "PATCH"},
		{&BeegoInput{Context: &Context{Request: &http.Request{Method: "HEAD"}}}, "HEAD"},
		{&BeegoInput{Context: &Context{Request: &http.Request{Method: "OPTIONS"}}}, "OPTIONS"},
	}

	for _, tt := range tests {
		t.Run(tt.want, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			if got := tt.input.Method(); got != tt.want {
				t.Errorf("BeegoInput.Method() = %v, want %v", got, tt.want)
			}
		})
	}
}
