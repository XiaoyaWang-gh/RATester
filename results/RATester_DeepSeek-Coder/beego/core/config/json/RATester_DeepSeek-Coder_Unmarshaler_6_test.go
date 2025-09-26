package json

import (
	"fmt"
	"testing"
)

func TestUnmarshaler_6(t *testing.T) {
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
			name:   "TestUnmarshaler_Success",
			prefix: "test",
			obj: testStruct{
				Name: "test",
				Age:  20,
			},
			wantErr: false,
		},
		{
			name:   "TestUnmarshaler_Fail",
			prefix: "test",
			obj: testStruct{
				Name: "test",
				Age:  20,
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

			c := &JSONConfigContainer{
				data: map[string]interface{}{
					"test": map[string]interface{}{
						"name": "test",
						"age":  20,
					},
				},
			}
			if err := c.Unmarshaler(tt.prefix, &tt.obj); (err != nil) != tt.wantErr {
				t.Errorf("Unmarshaler() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
