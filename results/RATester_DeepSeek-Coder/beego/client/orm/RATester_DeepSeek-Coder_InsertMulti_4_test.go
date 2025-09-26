package orm

import (
	"fmt"
	"testing"
)

func TestInsertMulti_4(t *testing.T) {
	type TestStruct struct {
		ID   int
		Name string
	}

	testCases := []struct {
		name     string
		bulk     int
		mds      interface{}
		expected int64
	}{
		{
			name:     "Test case 1",
			bulk:     10,
			mds:      []TestStruct{{ID: 1, Name: "Test1"}, {ID: 2, Name: "Test2"}},
			expected: 2,
		},
		{
			name:     "Test case 2",
			bulk:     0,
			mds:      []TestStruct{{ID: 1, Name: "Test1"}, {ID: 2, Name: "Test2"}},
			expected: 2,
		},
		{
			name:     "Test case 3",
			bulk:     10,
			mds:      []TestStruct{},
			expected: 0,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			orm := &DoNothingOrm{}
			result, err := orm.InsertMulti(tc.bulk, tc.mds)
			if err != nil {
				t.Errorf("Expected no error, got %v", err)
			}
			if result != tc.expected {
				t.Errorf("Expected %v, got %v", tc.expected, result)
			}
		})
	}
}
