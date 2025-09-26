package web

import (
	"fmt"
	"reflect"
	"testing"
)

func TestbasicKind_1(t *testing.T) {
	tests := []struct {
		name    string
		v       reflect.Value
		want    kind
		wantErr bool
	}{
		{
			name: "Test bool",
			v:    reflect.ValueOf(true),
			want: boolKind,
		},
		{
			name: "Test int",
			v:    reflect.ValueOf(1),
			want: intKind,
		},
		{
			name: "Test uint",
			v:    reflect.ValueOf(uint(1)),
			want: uintKind,
		},
		{
			name: "Test float",
			v:    reflect.ValueOf(float32(1.0)),
			want: floatKind,
		},
		{
			name: "Test complex",
			v:    reflect.ValueOf(complex(1.0, 1.0)),
			want: complexKind,
		},
		{
			name: "Test string",
			v:    reflect.ValueOf("test"),
			want: stringKind,
		},
		{
			name:    "Test invalid type",
			v:       reflect.ValueOf([]int{1, 2, 3}),
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

			got, err := basicKind(tt.v)
			if (err != nil) != tt.wantErr {
				t.Errorf("basicKind() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("basicKind() = %v, want %v", got, tt.want)
			}
		})
	}
}
