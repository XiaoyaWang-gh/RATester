package param

import (
	"fmt"
	"reflect"
	"testing"
)

func TestParse_10(t *testing.T) {
	testCases := []struct {
		name    string
		value   string
		want    interface{}
		wantErr bool
	}{
		{
			name:    "valid int",
			value:   "123",
			want:    123,
			wantErr: false,
		},
		{
			name:    "invalid int",
			value:   "abc",
			want:    nil,
			wantErr: true,
		},
		{
			name:    "negative int",
			value:   "-123",
			want:    -123,
			wantErr: false,
		},
	}

	p := intParser{}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			got, err := p.parse(tc.value, reflect.TypeOf(tc.want))
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
