package context

import (
	"fmt"
	"net/url"
	"reflect"
	"testing"
)

func TestBindMap_1(t *testing.T) {
	type testStruct struct {
		Name string
		Age  int
	}

	testCases := []struct {
		name   string
		params *url.Values
		key    string
		typ    reflect.Type
		want   reflect.Value
	}{
		{
			name: "Test Case 1",
			params: &url.Values{
				"Name[0]": []string{"John"},
				"Age[0]":  []string{"30"},
			},
			key: "Name",
			typ: reflect.TypeOf(map[string]testStruct{}),
			want: reflect.ValueOf(map[string]testStruct{
				"0": {Name: "John", Age: 30},
			}),
		},
		{
			name: "Test Case 2",
			params: &url.Values{
				"Name[0]": []string{"Jane"},
				"Age[0]":  []string{"25"},
			},
			key: "Name",
			typ: reflect.TypeOf(map[string]testStruct{}),
			want: reflect.ValueOf(map[string]testStruct{
				"0": {Name: "Jane", Age: 25},
			}),
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			input := &BeegoInput{}
			got := input.bindMap(tc.params, tc.key, tc.typ)

			if !reflect.DeepEqual(got, tc.want.Interface()) {
				t.Errorf("Expected %v, got %v", tc.want.Interface(), got)
			}
		})
	}
}
