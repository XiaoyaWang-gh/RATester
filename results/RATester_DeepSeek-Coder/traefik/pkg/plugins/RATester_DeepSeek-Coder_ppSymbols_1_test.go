package plugins

import (
	"fmt"
	"reflect"
	"testing"
)

func TestPpSymbols_1(t *testing.T) {
	testCases := []struct {
		name string
		want map[string]map[string]reflect.Value
	}{
		{
			name: "Test case 1",
			want: map[string]map[string]reflect.Value{
				"github.com/traefik/traefik/v3/pkg/plugins/plugins": {
					"PP":  reflect.ValueOf((*PP)(nil)),
					"_PP": reflect.ValueOf((*_PP)(nil)),
				},
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			got := ppSymbols()
			if !reflect.DeepEqual(got, tc.want) {
				t.Errorf("ppSymbols() = %v, want %v", got, tc.want)
			}
		})
	}
}
