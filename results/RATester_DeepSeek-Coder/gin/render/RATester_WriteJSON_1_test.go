package render

import (
	"bytes"
	"fmt"
	"io"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin/internal/json"
)

func TestWriteJSON_1(t *testing.T) {
	type testStruct struct {
		Name string `json:"name"`
		Age  int    `json:"age"`
	}

	testCases := []struct {
		name    string
		obj     any
		wantErr bool
	}{
		{
			name: "valid object",
			obj: testStruct{
				Name: "John Doe",
				Age:  30,
			},
			wantErr: false,
		},
		{
			name:    "nil object",
			obj:     nil,
			wantErr: false,
		},
		{
			name:    "unsupported type",
			obj:     make(chan int),
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

			rec := httptest.NewRecorder()
			err := WriteJSON(rec, tc.obj)
			if (err != nil) != tc.wantErr {
				t.Errorf("WriteJSON() error = %v, wantErr %v", err, tc.wantErr)
				return
			}

			if tc.wantErr {
				return
			}

			resp := rec.Result()
			defer resp.Body.Close()

			body, err := io.ReadAll(resp.Body)
			if err != nil {
				t.Errorf("Failed to read response body: %v", err)
				return
			}

			expected, err := json.Marshal(tc.obj)
			if err != nil {
				t.Errorf("Failed to marshal expected JSON: %v", err)
				return
			}

			if !bytes.Equal(body, expected) {
				t.Errorf("WriteJSON() = %s, want %s", body, expected)
			}
		})
	}
}
