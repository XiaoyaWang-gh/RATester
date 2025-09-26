package hugofs

import (
	"fmt"
	"reflect"
	"testing"

	radix "github.com/armon/go-radix"
)

func TestdoDoStat_1(t *testing.T) {
	fs := &RootMappingFs{
		id:            "test",
		rootMapToReal: radix.New(),
		realMapToRoot: radix.New(),
	}

	testCases := []struct {
		name     string
		setup    func()
		expected []FileMetaInfo
		err      error
	}{
		{
			name: "test case 1",
			setup: func() {
				// setup test case 1
			},
			expected: []FileMetaInfo{
				// expected result for test case 1
			},
			err: nil,
		},
		// add more test cases as needed
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			tc.setup()
			result, err := fs.doDoStat("test")
			if err != tc.err {
				t.Errorf("Expected error %v, but got %v", tc.err, err)
			}
			if !reflect.DeepEqual(result, tc.expected) {
				t.Errorf("Expected result %v, but got %v", tc.expected, result)
			}
		})
	}
}
