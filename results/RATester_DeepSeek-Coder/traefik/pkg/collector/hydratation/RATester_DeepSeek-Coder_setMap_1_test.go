package hydratation

import (
	"fmt"
	"reflect"
	"testing"
)

func TestSetMap_1(t *testing.T) {
	testCases := []struct {
		name    string
		field   reflect.Value
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			err := setMap(tc.field)
			if (err != nil) != tc.wantErr {
				t.Errorf("setMap() error = %v, wantErr %v", err, tc.wantErr)
			}
		})
	}
}
