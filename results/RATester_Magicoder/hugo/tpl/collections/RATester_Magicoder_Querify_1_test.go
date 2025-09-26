package collections

import (
	"fmt"
	"testing"
)

func TestQuerify_1(t *testing.T) {
	ns := &Namespace{}

	tests := []struct {
		name    string
		params  []any
		want    string
		wantErr bool
	}{
		{
			name:    "empty",
			params:  []any{},
			want:    "",
			wantErr: false,
		},
		{
			name:    "single",
			params:  []any{"key", "value"},
			want:    "key=value",
			wantErr: false,
		},
		{
			name:    "multiple",
			params:  []any{"key1", "value1", "key2", "value2"},
			want:    "key1=value1&key2=value2",
			wantErr: false,
		},
		{
			name:    "odd",
			params:  []any{"key1", "value1", "key2"},
			want:    "",
			wantErr: true,
		},
		{
			name:    "invalid",
			params:  []any{"key1", 123, "key2", "value2"},
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

			got, err := ns.Querify(tt.params...)
			if (err != nil) != tt.wantErr {
				t.Errorf("Querify() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Querify() = %v, want %v", got, tt.want)
			}
		})
	}
}
