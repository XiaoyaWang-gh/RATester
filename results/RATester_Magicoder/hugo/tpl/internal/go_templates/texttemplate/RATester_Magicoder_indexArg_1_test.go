package template

import (
	"fmt"
	"reflect"
	"testing"
)

func TestindexArg_1(t *testing.T) {
	tests := []struct {
		name    string
		index   reflect.Value
		cap     int
		want    int
		wantErr bool
	}{
		{
			name:    "Test with valid index",
			index:   reflect.ValueOf(1),
			cap:     10,
			want:    1,
			wantErr: false,
		},
		{
			name:    "Test with invalid index",
			index:   reflect.ValueOf(-1),
			cap:     10,
			want:    0,
			wantErr: true,
		},
		{
			name:    "Test with nil index",
			index:   reflect.ValueOf(nil),
			cap:     10,
			want:    0,
			wantErr: true,
		},
		{
			name:    "Test with index out of range",
			index:   reflect.ValueOf(11),
			cap:     10,
			want:    0,
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

			got, err := indexArg(tt.index, tt.cap)
			if (err != nil) != tt.wantErr {
				t.Errorf("indexArg() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("indexArg() = %v, want %v", got, tt.want)
			}
		})
	}
}
