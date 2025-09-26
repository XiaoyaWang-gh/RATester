package tableofcontents

import (
	"fmt"
	"reflect"
	"testing"
)

func TestWalk_4(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name      string
		fragments Fragments
		want      []string
	}{
		{
			name: "single heading",
			fragments: Fragments{
				Headings: Headings{
					{
						ID:    "single",
						Level: 1,
						Title: "Single",
					},
				},
			},
			want: []string{"single"},
		},
		{
			name: "multiple headings",
			fragments: Fragments{
				Headings: Headings{
					{
						ID:    "first",
						Level: 1,
						Title: "First",
						Headings: Headings{
							{
								ID:    "second",
								Level: 2,
								Title: "Second",
							},
						},
					},
					{
						ID:    "third",
						Level: 1,
						Title: "Third",
					},
				},
			},
			want: []string{"first", "second", "third"},
		},
	}

	for _, tt := range tests {
		tt := tt // NOTE: https://golang.org/doc/faq#closures_and_goroutines
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			t.Parallel()

			var got []string
			tt.fragments.walk(func(h *Heading) {
				got = append(got, h.ID)
			})

			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("got = %v, want = %v", got, tt.want)
			}
		})
	}
}
