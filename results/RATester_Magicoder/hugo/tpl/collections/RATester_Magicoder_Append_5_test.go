package collections

import (
	"fmt"
	"reflect"
	"testing"
)

func TestAppend_5(t *testing.T) {
	ns := &Namespace{}

	testCases := []struct {
		name    string
		args    []any
		want    any
		wantErr bool
	}{
		{
			name:    "Empty args",
			args:    []any{},
			want:    nil,
			wantErr: true,
		},
		{
			name:    "One arg",
			args:    []any{1},
			want:    nil,
			wantErr: true,
		},
		{
			name:    "Two args",
			args:    []any{1, 2},
			want:    []any{1, 2},
			wantErr: false,
		},
		{
			name:    "Three args",
			args:    []any{1, 2, 3},
			want:    []any{1, 2, 3},
			wantErr: false,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			got, err := ns.Append(tc.args...)
			if (err != nil) != tc.wantErr {
				t.Errorf("Append() error = %v, wantErr %v", err, tc.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tc.want) {
				t.Errorf("Append() = %v, want %v", got, tc.want)
			}
		})
	}
}
