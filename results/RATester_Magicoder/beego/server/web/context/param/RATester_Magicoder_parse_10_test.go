package param

import (
	"fmt"
	"reflect"
	"testing"
)

func Testparse_10(t *testing.T) {
	parser := intParser{}
	tests := []struct {
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
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			got, err := parser.parse(tt.value, reflect.TypeOf(tt.want))
			if (err != nil) != tt.wantErr {
				t.Errorf("parse() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("parse() = %v, want %v", got, tt.want)
			}
		})
	}
}
