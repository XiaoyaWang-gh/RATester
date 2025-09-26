package converter

import (
	"fmt"
	"reflect"
	"testing"
)

func TestBytes_1(t *testing.T) {
	type fields struct {
		b []byte
	}
	tests := []struct {
		name   string
		fields fields
		want   []byte
	}{
		{
			name: "Test Case 1",
			fields: fields{
				b: []byte("test"),
			},
			want: []byte("test"),
		},
		{
			name: "Test Case 2",
			fields: fields{
				b: []byte(""),
			},
			want: []byte(""),
		},
		{
			name: "Test Case 3",
			fields: fields{
				b: nil,
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

			b := Bytes(tt.fields.b)
			if got := b.Bytes(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Bytes() = %v, want %v", got, tt.want)
			}
		})
	}
}
