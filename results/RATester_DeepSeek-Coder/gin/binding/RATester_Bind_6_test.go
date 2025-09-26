package binding

import (
	"fmt"
	"net/http"
	"testing"
)

func TestBind_6(t *testing.T) {
	type testStruct struct {
		Field1 string `query:"field1"`
		Field2 int    `query:"field2"`
	}

	testCases := []struct {
		name       string
		query      string
		wantErr    bool
		wantObject testStruct
	}{
		{
			name:    "valid query",
			query:   "field1=value1&field2=2",
			wantErr: false,
			wantObject: testStruct{
				Field1: "value1",
				Field2: 2,
			},
		},
		{
			name:    "invalid query",
			query:   "field1=value1&field2=not_an_int",
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

			req, _ := http.NewRequest("GET", "http://test.com?"+tc.query, nil)
			obj := &testStruct{}
			err := queryBinding{}.Bind(req, obj)

			if (err != nil) != tc.wantErr {
				t.Errorf("Bind() error = %v, wantErr %v", err, tc.wantErr)
				return
			}

			if !tc.wantErr && *obj != tc.wantObject {
				t.Errorf("Bind() = %v, want %v", *obj, tc.wantObject)
			}
		})
	}
}
