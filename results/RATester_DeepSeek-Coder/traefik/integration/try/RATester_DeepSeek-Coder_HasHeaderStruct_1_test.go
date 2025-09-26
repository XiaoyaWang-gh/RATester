package try

import (
	"fmt"
	"net/http"
	"reflect"
	"testing"
)

func TestHasHeaderStruct_1(t *testing.T) {
	type args struct {
		header http.Header
	}
	tests := []struct {
		name string
		args args
		want ResponseCondition
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			if got := HasHeaderStruct(tt.args.header); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("HasHeaderStruct() = %v, want %v", got, tt.want)
			}
		})
	}
}
