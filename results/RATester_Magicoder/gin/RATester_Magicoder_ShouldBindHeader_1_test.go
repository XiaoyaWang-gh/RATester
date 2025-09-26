package gin

import (
	"errors"
	"fmt"
	"net/http"
	"testing"
)

func TestShouldBindHeader_1(t *testing.T) {
	type testStruct struct {
		Field1 string `header:"Field1"`
		Field2 string `header:"Field2"`
	}

	testCases := []struct {
		name   string
		header http.Header
		obj    testStruct
		want   error
	}{
		{
			name: "valid header",
			header: http.Header{
				"Field1": []string{"value1"},
				"Field2": []string{"value2"},
			},
			obj:  testStruct{},
			want: nil,
		},
		{
			name: "invalid header",
			header: http.Header{
				"Field1": []string{"value1"},
			},
			obj:  testStruct{},
			want: errors.New("missing required header field"),
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			req, err := http.NewRequest("GET", "/", nil)
			if err != nil {
				t.Fatal(err)
			}
			req.Header = tc.header

			c := &Context{Request: req}
			err = c.ShouldBindHeader(&tc.obj)

			if err != nil && tc.want != nil {
				if err.Error() != tc.want.Error() {
					t.Errorf("got error %v, want %v", err, tc.want)
				}
			} else if err == nil && tc.want != nil {
				t.Errorf("expected error, got nil")
			} else if err != nil && tc.want == nil {
				t.Errorf("unexpected error: %v", err)
			}
		})
	}
}
