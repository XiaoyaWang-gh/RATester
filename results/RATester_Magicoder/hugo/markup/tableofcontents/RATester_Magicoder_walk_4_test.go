package tableofcontents

import (
	"fmt"
	"reflect"
	"testing"
)

func Testwalk_4(t *testing.T) {
	testCases := []struct {
		name      string
		fragments Fragments
		expected  []string
	}{
		{
			name: "Test case 1",
			fragments: Fragments{
				Headings: Headings{
					{
						ID:    "id1",
						Level: 1,
						Title: "Title1",
						Headings: Headings{
							{
								ID:    "id2",
								Level: 2,
								Title: "Title2",
								Headings: Headings{
									{
										ID:    "id3",
										Level: 3,
										Title: "Title3",
									},
								},
							},
						},
					},
				},
			},
			expected: []string{"id1", "id2", "id3"},
		},
		// Add more test cases as needed
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			var ids []string
			tc.fragments.walk(func(h *Heading) {
				ids = append(ids, h.ID)
			})

			if !reflect.DeepEqual(ids, tc.expected) {
				t.Errorf("Expected: %v, Got: %v", tc.expected, ids)
			}
		})
	}
}
