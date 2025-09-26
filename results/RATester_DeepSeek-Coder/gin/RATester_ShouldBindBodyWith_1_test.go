package gin

import (
	"fmt"
	"net/http"
	"reflect"
	"strings"
	"testing"

	"github.com/gin-gonic/gin/binding"
)

func TestShouldBindBodyWith_1(t *testing.T) {
	type testStruct struct {
		Field string `json:"field"`
	}

	testCases := []struct {
		name     string
		body     string
		expected testStruct
		wantErr  bool
	}{
		{
			name:     "valid JSON",
			body:     `{"field": "value"}`,
			expected: testStruct{Field: "value"},
			wantErr:  false,
		},
		{
			name:    "invalid JSON",
			body:    `{"field": "value",}`,
			wantErr: true,
		},
		{
			name:    "empty body",
			body:    "",
			wantErr: true,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			req, err := http.NewRequest("POST", "/", strings.NewReader(tc.body))
			if err != nil {
				t.Fatal(err)
			}

			ctx := &Context{Request: req}
			obj := &testStruct{}

			err = ctx.ShouldBindBodyWith(obj, binding.JSON)
			if (err != nil) != tc.wantErr {
				t.Errorf("ShouldBindBodyWith() error = %v, wantErr %v", err, tc.wantErr)
			}

			if !tc.wantErr {
				if !reflect.DeepEqual(*obj, tc.expected) {
					t.Errorf("ShouldBindBodyWith() = %v, want %v", *obj, tc.expected)
				}
			}
		})
	}
}
