package json

import (
	"fmt"
	"testing"
)

func TestInt_7(t *testing.T) {
	t.Run("TestInt", func(t *testing.T) {

		defer func() {
			if r := recover(); r != nil {
				fmt.Println("Recovered in main", r)
			}
		}()

		container := &JSONConfigContainer{
			data: map[string]interface{}{
				"key1": float64(123),
				"key2": "456",
				"key3": "not a number",
				"key4": nil,
			},
		}

		tests := []struct {
			name    string
			key     string
			want    int
			wantErr bool
		}{
			{"TestInt_1", "key1", 123, false},
			{"TestInt_2", "key2", 456, false},
			{"TestInt_3", "key3", 0, true},
			{"TestInt_4", "key4", 0, true},
			{"TestInt_5", "key5", 0, true},
		}

		for _, tt := range tests {
			t.Run(tt.name, func(t *testing.T) {

				defer func() {
					if r := recover(); r != nil {
						fmt.Println("Recovered in main", r)
					}
				}()

				got, err := container.Int(tt.key)
				if (err != nil) != tt.wantErr {
					t.Errorf("Int() error = %v, wantErr %v", err, tt.wantErr)
					return
				}
				if got != tt.want {
					t.Errorf("Int() got = %v, want %v", got, tt.want)
				}
			})
		}
	})
}
