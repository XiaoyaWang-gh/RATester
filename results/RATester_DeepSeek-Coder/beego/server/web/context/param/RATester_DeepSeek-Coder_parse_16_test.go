package param

import (
	"fmt"
	"reflect"
	"testing"
)

func TestParse_16(t *testing.T) {
	testCases := []struct {
		name    string
		value   string
		toType  reflect.Type
		want    interface{}
		wantErr bool
	}{
		{
			name:    "valid bool",
			value:   "true",
			toType:  reflect.TypeOf(true),
			want:    true,
			wantErr: false,
		},
		{
			name:    "invalid bool",
			value:   "invalid",
			toType:  reflect.TypeOf(true),
			want:    false,
			wantErr: true,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			p := boolParser{}
			got, err := p.parse(tc.value, tc.toType)
			if (err != nil) != tc.wantErr {
				t.Errorf("parse() error = %v, wantErr %v", err, tc.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tc.want) {
				t.Errorf("parse() = %v, want %v", got, tc.want)
			}
		})
	}
}
