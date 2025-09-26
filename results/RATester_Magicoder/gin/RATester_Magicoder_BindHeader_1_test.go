package gin

import (
	"fmt"
	"net/http"
	"reflect"
	"testing"
)

func TestBindHeader_1(t *testing.T) {
	type testStruct struct {
		Field1 string `header:"Field1"`
		Field2 int    `header:"Field2"`
	}

	testCases := []struct {
		name    string
		header  http.Header
		want    testStruct
		wantErr bool
	}{
		{
			name: "valid header",
			header: http.Header{
				"Field1": []string{"value1"},
				"Field2": []string{"123"},
			},
			want: testStruct{
				Field1: "value1",
				Field2: 123,
			},
			wantErr: false,
		},
		{
			name: "invalid header",
			header: http.Header{
				"Field1": []string{"value1"},
				"Field2": []string{"abc"},
			},
			want:    testStruct{},
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

			req, err := http.NewRequest("GET", "/", nil)
			if err != nil {
				t.Fatal(err)
			}
			req.Header = tc.header

			ctx := &Context{Request: req}
			obj := testStruct{}

			err = ctx.BindHeader(&obj)
			if (err != nil) != tc.wantErr {
				t.Errorf("BindHeader() error = %v, wantErr %v", err, tc.wantErr)
				return
			}

			if !reflect.DeepEqual(obj, tc.want) {
				t.Errorf("BindHeader() = %v, want %v", obj, tc.want)
			}
		})
	}
}
