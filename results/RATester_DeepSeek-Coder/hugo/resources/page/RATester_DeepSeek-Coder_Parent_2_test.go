package page

import (
	"fmt"
	"reflect"
	"testing"
)

func TestParent_2(t *testing.T) {
	type testStruct struct {
		nopPage nopPage
	}

	tests := []struct {
		name string
		ts   testStruct
		want Page
	}{
		{
			name: "Test Parent",
			ts: testStruct{
				nopPage: nopPage(1),
			},
			want: nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			if got := tt.ts.nopPage.Parent(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Parent() = %v, want %v", got, tt.want)
			}
		})
	}
}
