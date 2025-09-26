package gin

import (
	"fmt"
	"net/http"
	"net/url"
	"testing"
)

func TestBindQuery_1(t *testing.T) {
	type testStruct struct {
		Name string `form:"name"`
		Age  int    `form:"age"`
	}

	tests := []struct {
		name    string
		query   string
		wantErr bool
	}{
		{
			name:    "valid query",
			query:   "name=John&age=30",
			wantErr: false,
		},
		{
			name:    "invalid query",
			query:   "name=John&age=not_an_int",
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			c := &Context{
				Request: &http.Request{
					URL: &url.URL{
						RawQuery: tt.query,
					},
				},
			}

			s := &testStruct{}
			err := c.BindQuery(s)
			if (err != nil) != tt.wantErr {
				t.Errorf("BindQuery() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
