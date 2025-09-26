package gin

import (
	"fmt"
	"net/url"
	"reflect"
	"testing"
)

func TestPostFormMap_1(t *testing.T) {
	type testCase struct {
		name string
		key  string
		want map[string]string
	}

	tests := []testCase{
		{
			name: "Test case 1",
			key:  "key1",
			want: map[string]string{"key1": "value1"},
		},
		{
			name: "Test case 2",
			key:  "key2",
			want: map[string]string{"key2": "value2"},
		},
		// Add more test cases as needed
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			c := &Context{
				formCache: url.Values{
					tt.key: []string{tt.want[tt.key]},
				},
			}
			if got := c.PostFormMap(tt.key); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("PostFormMap() = %v, want %v", got, tt.want)
			}
		})
	}
}
