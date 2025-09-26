package xml

import (
	"fmt"
	"testing"
)

func TestUnmarshaler_9(t *testing.T) {
	t.Parallel()

	type testStruct struct {
		Name string `mapstructure:"name"`
		Age  int    `mapstructure:"age"`
	}

	tests := []struct {
		name    string
		prefix  string
		obj     testStruct
		wantErr bool
	}{
		{
			name:   "success",
			prefix: "test",
			obj: testStruct{
				Name: "test",
				Age:  30,
			},
			wantErr: false,
		},
		{
			name:   "failure",
			prefix: "test",
			obj: testStruct{
				Name: "test",
				Age:  30,
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

			c := &ConfigContainer{
				data: map[string]interface{}{
					"test.name": "test",
					"test.age":  30,
				},
			}
			if err := c.Unmarshaler(tt.prefix, &tt.obj); (err != nil) != tt.wantErr {
				t.Errorf("Unmarshaler() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
