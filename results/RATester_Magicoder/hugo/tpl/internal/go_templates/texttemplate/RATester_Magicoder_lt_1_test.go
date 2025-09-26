package template

import (
	"fmt"
	"reflect"
	"testing"
)

func Testlt_1(t *testing.T) {
	tests := []struct {
		name    string
		arg1    reflect.Value
		arg2    reflect.Value
		want    bool
		wantErr bool
	}{
		{
			name: "Test case 1",
			arg1: reflect.ValueOf(1),
			arg2: reflect.ValueOf(2),
			want: true,
		},
		{
			name: "Test case 2",
			arg1: reflect.ValueOf(3),
			arg2: reflect.ValueOf(2),
			want: false,
		},
		{
			name: "Test case 3",
			arg1: reflect.ValueOf(4),
			arg2: reflect.ValueOf(4),
			want: false,
		},
		{
			name: "Test case 4",
			arg1: reflect.ValueOf(5),
			arg2: reflect.ValueOf(6),
			want: true,
		},
		{
			name: "Test case 5",
			arg1: reflect.ValueOf(7),
			arg2: reflect.ValueOf(8),
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			got, err := lt(tt.arg1, tt.arg2)
			if (err != nil) != tt.wantErr {
				t.Errorf("lt() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("lt() = %v, want %v", got, tt.want)
			}
		})
	}
}
