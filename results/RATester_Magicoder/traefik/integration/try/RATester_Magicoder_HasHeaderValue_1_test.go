package try

import (
	"errors"
	"fmt"
	"net/http"
	"reflect"
	"testing"
)

func TestHasHeaderValue_1(t *testing.T) {
	type args struct {
		header     string
		value      string
		exactMatch bool
	}
	tests := []struct {
		name string
		args args
		want ResponseCondition
	}{
		{
			name: "Test case 1",
			args: args{
				header:     "Content-Type",
				value:      "application/json",
				exactMatch: true,
			},
			want: func(res *http.Response) error {
				if _, ok := res.Header["Content-Type"]; !ok {
					return errors.New("response doesn't contain header: Content-Type")
				}

				matchFound := false
				for _, hdr := range res.Header["Content-Type"] {
					if "application/json" != hdr {
						return fmt.Errorf("got header Content-Type with value %s, wanted application/json", hdr)
					}
					if "application/json" == hdr {
						matchFound = true
					}
				}

				if !matchFound {
					return fmt.Errorf("response doesn't contain header Content-Type with value application/json")
				}
				return nil
			},
		},
		// Add more test cases as needed
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			if got := HasHeaderValue(tt.args.header, tt.args.value, tt.args.exactMatch); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("HasHeaderValue() = %v, want %v", got, tt.want)
			}
		})
	}
}
