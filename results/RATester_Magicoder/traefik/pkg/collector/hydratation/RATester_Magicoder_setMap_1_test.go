package hydratation

import (
	"fmt"
	"reflect"
	"testing"
)

func TestSetMap_1(t *testing.T) {
	type testStruct struct {
		field reflect.Value
	}

	tests := []struct {
		name    string
		ts      testStruct
		wantErr bool
	}{
		{
			name: "Test case 1",
			ts: testStruct{
				field: reflect.ValueOf(map[string]int{}),
			},
			wantErr: false,
		},
		{
			name: "Test case 2",
			ts: testStruct{
				field: reflect.ValueOf(nil),
			},
			wantErr: true,
		},
		{
			name: "Test case 3",
			ts: testStruct{
				field: reflect.ValueOf(map[string]int{}),
			},
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			if err := setMap(tt.ts.field); (err != nil) != tt.wantErr {
				t.Errorf("setMap() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
