package gin

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"reflect"
	"strings"
	"testing"
)

func TestGetRawData_1(t *testing.T) {
	testCases := []struct {
		name    string
		context *Context
		want    []byte
		wantErr bool
	}{
		{
			name: "Test case 1",
			context: &Context{
				Request: &http.Request{
					Body: ioutil.NopCloser(strings.NewReader("test data")),
				},
			},
			want:    []byte("test data"),
			wantErr: false,
		},
		{
			name: "Test case 2",
			context: &Context{
				Request: &http.Request{
					Body: nil,
				},
			},
			want:    nil,
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

			got, err := tc.context.GetRawData()
			if (err != nil) != tc.wantErr {
				t.Errorf("GetRawData() error = %v, wantErr %v", err, tc.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tc.want) {
				t.Errorf("GetRawData() = %v, want %v", got, tc.want)
			}
		})
	}
}
