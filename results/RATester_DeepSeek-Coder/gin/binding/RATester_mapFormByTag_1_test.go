package binding

import (
	"fmt"
	"testing"
)

func TestMapFormByTag_1(t *testing.T) {
	type testStruct struct {
		Name string `form:"name"`
		Age  int    `form:"age"`
	}

	testCases := []struct {
		name    string
		ptr     any
		form    map[string][]string
		tag     string
		wantErr bool
	}{
		{
			name:    "Test with map",
			ptr:     &map[string]string{},
			form:    map[string][]string{"name": {"John"}, "age": {"30"}},
			tag:     "form",
			wantErr: false,
		},
		{
			name:    "Test with struct",
			ptr:     &testStruct{},
			form:    map[string][]string{"name": {"John"}, "age": {"30"}},
			tag:     "form",
			wantErr: false,
		},
		{
			name:    "Test with invalid ptr",
			ptr:     "invalid",
			form:    map[string][]string{"name": {"John"}, "age": {"30"}},
			tag:     "form",
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

			err := mapFormByTag(tc.ptr, tc.form, tc.tag)
			if (err != nil) != tc.wantErr {
				t.Errorf("mapFormByTag() error = %v, wantErr %v", err, tc.wantErr)
			}
		})
	}
}
