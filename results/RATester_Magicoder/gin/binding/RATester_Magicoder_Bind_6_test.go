package binding

import (
	"fmt"
	"net/http"
	"testing"
)

func TestBind_6(t *testing.T) {
	type testStruct struct {
		Field1 string `form:"field1"`
		Field2 int    `form:"field2"`
	}

	testCases := []struct {
		name    string
		query   string
		obj     testStruct
		wantErr bool
	}{
		{
			name:  "valid query",
			query: "field1=value1&field2=123",
			obj:   testStruct{},
		},
		{
			name:    "invalid query",
			query:   "field1=value1&field2=abc",
			obj:     testStruct{},
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

			req, err := http.NewRequest("GET", "http://test.com/?"+tc.query, nil)
			if err != nil {
				t.Fatal(err)
			}

			queryBinding := queryBinding{}
			err = queryBinding.Bind(req, &tc.obj)

			if (err != nil) != tc.wantErr {
				t.Errorf("Bind() error = %v, wantErr %v", err, tc.wantErr)
				return
			}
		})
	}
}
