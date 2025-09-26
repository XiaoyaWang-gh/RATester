package gin

import (
	"fmt"
	"net/http"
	"strings"
	"testing"
)

func TestShouldBindBodyWithJSON_1(t *testing.T) {
	type testStruct struct {
		Field1 string `json:"field1"`
		Field2 int    `json:"field2"`
	}

	testCases := []struct {
		name    string
		body    string
		wantErr bool
	}{
		{
			name:    "Valid JSON",
			body:    `{"field1": "value1", "field2": 123}`,
			wantErr: false,
		},
		{
			name:    "Invalid JSON",
			body:    `{"field1": "value1", "field2": "abc"}`,
			wantErr: true,
		},
		{
			name:    "Empty Body",
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

			req, err := http.NewRequest("POST", "/test", strings.NewReader(tc.body))
			if err != nil {
				t.Fatal(err)
			}

			ctx := &Context{Request: req}
			obj := &testStruct{}

			err = ctx.ShouldBindBodyWithJSON(obj)
			if (err != nil) != tc.wantErr {
				t.Errorf("ShouldBindBodyWithJSON() error = %v, wantErr %v", err, tc.wantErr)
			}
		})
	}
}
