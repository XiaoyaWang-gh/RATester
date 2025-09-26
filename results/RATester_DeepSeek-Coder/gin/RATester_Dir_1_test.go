package gin

import (
	"fmt"
	"net/http"
	"reflect"
	"testing"
)

func TestDir_1(t *testing.T) {
	testCases := []struct {
		name           string
		root           string
		listDirectory  bool
		expectedResult http.FileSystem
	}{
		{
			name:           "Test Case 1",
			root:           "/test",
			listDirectory:  true,
			expectedResult: http.Dir("/test"),
		},
		{
			name:           "Test Case 2",
			root:           "/test",
			listDirectory:  false,
			expectedResult: &OnlyFilesFS{FileSystem: http.Dir("/test")},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			result := Dir(tc.root, tc.listDirectory)
			if !reflect.DeepEqual(result, tc.expectedResult) {
				t.Errorf("Expected %v, got %v", tc.expectedResult, result)
			}
		})
	}
}
