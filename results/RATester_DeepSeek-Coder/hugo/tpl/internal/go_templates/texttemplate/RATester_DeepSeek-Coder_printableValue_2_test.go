package template

import (
	"fmt"
	"reflect"
	"testing"
)

func TestPrintableValue_2(t *testing.T) {
	tests := []struct {
		name string
		v    reflect.Value
		want any
		ok   bool
	}{
		{
			name: "Test case 1",
			v:    reflect.ValueOf("test"),
			want: "test",
			ok:   true,
		},
		{
			name: "Test case 2",
			v:    reflect.ValueOf(123),
			want: 123,
			ok:   true,
		},
		{
			name: "Test case 3",
			v:    reflect.ValueOf(nil),
			want: "<no value>",
			ok:   true,
		},
		{
			name: "Test case 4",
			v:    reflect.ValueOf(make(chan int)),
			want: nil,
			ok:   false,
		},
		{
			name: "Test case 5",
			v:    reflect.ValueOf(func() {}),
			want: nil,
			ok:   false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			got, ok := printableValue(tt.v)
			if !reflect.DeepEqual(got, tt.want) || ok != tt.ok {
				t.Errorf("printableValue() = %v, %v; want %v, %v", got, ok, tt.want, tt.ok)
			}
		})
	}
}
