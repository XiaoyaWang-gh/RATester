package yaml

import (
	"fmt"
	"testing"
)

func TestInt64_1(t *testing.T) {
	type testCase struct {
		name    string
		key     string
		value   interface{}
		want    int64
		wantErr bool
	}

	tests := []testCase{
		{
			name:    "int",
			key:     "key1",
			value:   10,
			want:    10,
			wantErr: false,
		},
		{
			name:    "int64",
			key:     "key2",
			value:   int64(20),
			want:    20,
			wantErr: false,
		},
		{
			name:    "string",
			key:     "key3",
			value:   "30",
			want:    0,
			wantErr: true,
		},
		{
			name:    "float64",
			key:     "key4",
			value:   40.0,
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

			c := &ConfigContainer{
				data: map[string]interface{}{
					tt.key: tt.value,
				},
			}
			got, err := c.Int64(tt.key)
			if (err != nil) != tt.wantErr {
				t.Errorf("Int64() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Int64() = %v, want %v", got, tt.want)
			}
		})
	}
}
