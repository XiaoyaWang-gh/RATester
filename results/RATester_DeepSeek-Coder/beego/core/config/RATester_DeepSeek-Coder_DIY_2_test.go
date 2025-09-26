package config

import (
	"fmt"
	"reflect"
	"testing"
)

func TestDIY_2(t *testing.T) {
	t.Run("TestDIY", func(t *testing.T) {

		defer func() {
			if r := recover(); r != nil {
				fmt.Println("Recovered in main", r)
			}
		}()

		testDIY := &IniConfigContainer{
			data: map[string]map[string]string{
				"section1": {"key1": "value1", "key2": "value2"},
				"section2": {"key3": "value3", "key4": "value4"},
			},
		}

		tests := []struct {
			name    string
			key     string
			wantV   interface{}
			wantErr bool
		}{
			{"TestDIY_0", "section1:key1", "value1", false},
			{"TestDIY_1", "section1:key3", nil, true},
			{"TestDIY_2", "section2:key4", "value4", false},
			{"TestDIY_3", "section2:key2", "value2", false},
			{"TestDIY_4", "section3:key5", nil, true},
		}

		for _, tt := range tests {
			t.Run(tt.name, func(t *testing.T) {

				defer func() {
					if r := recover(); r != nil {
						fmt.Println("Recovered in main", r)
					}
				}()

				gotV, err := testDIY.DIY(tt.key)
				if (err != nil) != tt.wantErr {
					t.Errorf("DIY() error = %v, wantErr %v", err, tt.wantErr)
					return
				}
				if !reflect.DeepEqual(gotV, tt.wantV) {
					t.Errorf("DIY() = %v, want %v", gotV, tt.wantV)
				}
			})
		}
	})
}
