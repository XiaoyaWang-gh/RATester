package collections

import (
	"fmt"
	"reflect"
	"testing"
)

func TestcollectIdentities_1(t *testing.T) {
	tests := []struct {
		name    string
		args    []any
		want    map[any]bool
		wantErr bool
	}{
		{
			name: "Test case 1",
			args: []any{[]any{1, 2, 3}, []any{4, 5, 6}},
			want: map[any]bool{1: true, 2: true, 3: true, 4: true, 5: true, 6: true},
		},
		{
			name:    "Test case 2",
			args:    []any{[]any{1, 2, 3}, "not a slice or array"},
			wantErr: true,
		},
		{
			name:    "Test case 3",
			args:    []any{[]any{1, 2, 3}, []any{4, "not comparable", 6}},
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

			got, err := collectIdentities(tt.args...)
			if (err != nil) != tt.wantErr {
				t.Errorf("collectIdentities() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("collectIdentities() = %v, want %v", got, tt.want)
			}
		})
	}
}
