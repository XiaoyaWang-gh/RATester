package context

import (
	"fmt"
	"net/http"
	"strings"
	"testing"
)

func TestBind_1(t *testing.T) {
	type testStruct struct {
		Name string `form:"name"`
	}

	testCases := []struct {
		name       string
		content    string
		contentype string
		wantErr    bool
	}{
		{"JSON", `{"name":"test"}`, "application/json", false},
		{"XML", `<testStruct><Name>test</Name></testStruct>`, "application/xml", false},
		{"Form", "name=test", "application/x-www-form-urlencoded", false},
		{"Protobuf", "", "application/protobuf", true}, // No protobuf data provided
		{"YAML", "name: test", "application/x-yaml", false},
		{"Unsupported", "", "unsupported/type", true},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			req, err := http.NewRequest("POST", "/", strings.NewReader(tc.content))
			if err != nil {
				t.Fatal(err)
			}
			req.Header.Set("Content-Type", tc.contentype)

			ctx := &Context{Request: req}
			obj := &testStruct{}

			err = ctx.Bind(obj)
			if (err != nil) != tc.wantErr {
				t.Errorf("Bind() error = %v, wantErr %v", err, tc.wantErr)
			}
		})
	}
}
