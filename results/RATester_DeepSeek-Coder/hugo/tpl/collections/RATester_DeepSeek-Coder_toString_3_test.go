package collections

import (
	"fmt"
	"reflect"
	"testing"
)

func TestToString_3(t *testing.T) {
	tests := []struct {
		name    string
		v       reflect.Value
		want    string
		wantErr bool
	}{
		{
			name:    "Test case 1",
			v:       reflect.ValueOf("test"),
			want:    "test",
			wantErr: false,
		},
		{
			name:    "Test case 2",
			v:       reflect.ValueOf(123),
			want:    "",
			wantErr: true,
		},
		{
			name:    "Test case 3",
			v:       reflect.ValueOf(nil),
			want:    "",
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

			got, err := toString(tt.v)
			if (err != nil) != tt.wantErr {
				t.Errorf("toString() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("toString() = %v, want %v", got, tt.want)
			}
		})
	}
}
