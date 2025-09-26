package context

import (
	"fmt"
	"net/url"
	"reflect"
	"testing"
)

func TestFormValue_1(t *testing.T) {
	type testStruct struct {
		Field string `default:"defaultValue"`
	}

	testCases := []struct {
		name     string
		tag      string
		form     url.Values
		expected string
		ok       bool
	}{
		{
			name:     "Test with no form values and default value",
			tag:      "Field",
			form:     url.Values{},
			expected: "defaultValue",
			ok:       true,
		},
		{
			name:     "Test with one form value",
			tag:      "Field",
			form:     url.Values{"Field": []string{"formValue"}},
			expected: "formValue",
			ok:       true,
		},
		{
			name:     "Test with multiple form values",
			tag:      "Field",
			form:     url.Values{"Field": []string{"formValue1", "formValue2"}},
			expected: "formValue1",
			ok:       true,
		},
		{
			name:     "Test with no form values and no default value",
			tag:      "Field",
			form:     url.Values{},
			expected: "",
			ok:       false,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			fieldT := reflect.StructField{
				Name: "Field",
				Type: reflect.TypeOf(""),
				Tag:  reflect.StructTag(`default:"defaultValue"`),
			}
			actual, ok := formValue(tc.tag, tc.form, fieldT)
			if actual != tc.expected {
				t.Errorf("Expected %v, got %v", tc.expected, actual)
			}
			if ok != tc.ok {
				t.Errorf("Expected %v, got %v", tc.ok, ok)
			}
		})
	}
}
