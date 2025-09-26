package identity

import (
	"fmt"
	"reflect"
	"testing"
)

func TestWalkIdentitiesDeep_1(t *testing.T) {
	type testIdentity struct {
		id string
	}

	tests := []struct {
		name string
		v    any
		want []string
	}{
		{
			name: "Test with a map",
			v: map[string]any{
				"key1": "value1",
				"key2": testIdentity{id: "id1"},
			},
			want: []string{"value1", "id1"},
		},
		{
			name: "Test with a slice",
			v: []any{
				"value2",
				testIdentity{id: "id2"},
			},
			want: []string{"value2", "id2"},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			got := make([]string, 0)
			WalkIdentitiesDeep(tt.v, func(level int, id Identity) bool {
				if id != nil {
					got = append(got, id.IdentifierBase())
				} else {
					got = append(got, fmt.Sprintf("%v", id))
				}
				return true
			})

			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("WalkIdentitiesDeep() = %v, want %v", got, tt.want)
			}
		})
	}
}
