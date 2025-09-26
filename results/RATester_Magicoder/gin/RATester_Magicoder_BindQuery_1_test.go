package gin

import (
	"fmt"
	"net/http"
	"testing"
)

func TestBindQuery_1(t *testing.T) {
	type testStruct struct {
		Field1 string `form:"field1"`
		Field2 int    `form:"field2"`
	}

	testCases := []struct {
		name    string
		query   string
		want    testStruct
		wantErr bool
	}{
		{
			name:  "valid query",
			query: "field1=value1&field2=123",
			want:  testStruct{Field1: "value1", Field2: 123},
		},
		{
			name:    "invalid query",
			query:   "field1=value1&field2=abc",
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

			req, err := http.NewRequest("GET", "/?"+tc.query, nil)
			if err != nil {
				t.Fatal(err)
			}

			ctx := &Context{Request: req}
			obj := testStruct{}

			err = ctx.BindQuery(&obj)
			if (err != nil) != tc.wantErr {
				t.Errorf("BindQuery() error = %v, wantErr %v", err, tc.wantErr)
				return
			}

			if !tc.wantErr {
				if obj.Field1 != tc.want.Field1 || obj.Field2 != tc.want.Field2 {
					t.Errorf("BindQuery() = %v, want %v", obj, tc.want)
				}
			}
		})
	}
}
