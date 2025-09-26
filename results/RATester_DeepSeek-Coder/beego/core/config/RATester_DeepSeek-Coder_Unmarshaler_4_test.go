package config

import (
	"fmt"
	"testing"
)

func TestUnmarshaler_4(t *testing.T) {
	type testStruct struct {
		Name string `json:"name"`
		Age  int    `json:"age"`
	}

	tests := []struct {
		name    string
		prefix  string
		obj     testStruct
		wantErr bool
	}{
		{
			name:   "test case 1",
			prefix: "",
			obj: testStruct{
				Name: "test",
				Age:  18,
			},
			wantErr: false,
		},
		{
			name:   "test case 2",
			prefix: "test",
			obj: testStruct{
				Name: "test",
				Age:  18,
			},
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

			c := &IniConfigContainer{
				data: map[string]map[string]string{
					"test": {
						"name": "test",
						"age":  "18",
					},
				},
			}
			if err := c.Unmarshaler(tt.prefix, &tt.obj); (err != nil) != tt.wantErr {
				t.Errorf("Unmarshaler() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
