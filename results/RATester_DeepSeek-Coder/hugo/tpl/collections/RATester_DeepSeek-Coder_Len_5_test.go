package collections

import (
	"fmt"
	"reflect"
	"testing"
)

func TestLen_5(t *testing.T) {
	testCases := []struct {
		name string
		p    pairList
		want int
	}{
		{
			name: "empty pairList",
			p:    pairList{},
			want: 0,
		},
		{
			name: "non-empty pairList",
			p: pairList{
				Pairs: []pair{
					{
						Key:   reflect.ValueOf("key1"),
						Value: reflect.ValueOf("value1"),
					},
					{
						Key:   reflect.ValueOf("key2"),
						Value: reflect.ValueOf("value2"),
					},
				},
			},
			want: 2,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			got := tc.p.Len()
			if got != tc.want {
				t.Errorf("got %v, want %v", got, tc.want)
			}
		})
	}
}
