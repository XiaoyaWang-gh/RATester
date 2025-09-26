package config

import (
	"fmt"
	"testing"
)

func TestUnmarshaler_4(t *testing.T) {
	type TestStruct struct {
		Key   string
		Value string
	}

	testCases := []struct {
		name    string
		prefix  string
		obj     interface{}
		wantErr bool
	}{
		{
			name:   "Test case 1",
			prefix: "",
			obj: &TestStruct{
				Key:   "test_key",
				Value: "test_value",
			},
			wantErr: false,
		},
		{
			name:   "Test case 2",
			prefix: "test_prefix",
			obj: &TestStruct{
				Key:   "test_key",
				Value: "test_value",
			},
			wantErr: true,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			iniConfig := &IniConfigContainer{
				data: map[string]map[string]string{
					"test_section": {
						"test_key": "test_value",
					},
				},
			}

			err := iniConfig.Unmarshaler(tc.prefix, tc.obj)
			if (err != nil) != tc.wantErr {
				t.Errorf("Unmarshaler() error = %v, wantErr %v", err, tc.wantErr)
				return
			}
		})
	}
}
