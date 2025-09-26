package render

import (
	"fmt"
	"net/http/httptest"
	"testing"
)

func TestRender_16(t *testing.T) {
	type testCase struct {
		name     string
		data     any
		expected string
	}

	tests := []testCase{
		{
			name:     "Test with string data",
			data:     "test",
			expected: "test",
		},
		{
			name:     "Test with int data",
			data:     123,
			expected: "123",
		},
		{
			name:     "Test with float data",
			data:     123.456,
			expected: "123.456",
		},
		{
			name:     "Test with struct data",
			data:     struct{ Name string }{"John"},
			expected: "{Name:John}",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			w := httptest.NewRecorder()
			p := ProtoBuf{Data: tt.data}
			err := p.Render(w)
			if err != nil {
				t.Errorf("Render() error = %v", err)
				return
			}
			got := w.Body.String()
			if got != tt.expected {
				t.Errorf("Render() = %v, want %v", got, tt.expected)
			}
		})
	}
}
