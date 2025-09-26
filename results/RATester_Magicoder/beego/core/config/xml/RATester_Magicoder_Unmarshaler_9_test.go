package xml

import (
	"fmt"
	"testing"
)

func TestUnmarshaler_9(t *testing.T) {
	type testStruct struct {
		Field1 string `mapstructure:"field1"`
		Field2 int    `mapstructure:"field2"`
	}

	testCases := []struct {
		name    string
		prefix  string
		obj     interface{}
		data    map[string]interface{}
		wantErr bool
	}{
		{
			name:   "Success",
			prefix: "test",
			obj:    &testStruct{},
			data: map[string]interface{}{
				"field1": "value1",
				"field2": 123,
			},
			wantErr: false,
		},
		{
			name:   "Error",
			prefix: "test",
			obj:    &testStruct{},
			data: map[string]interface{}{
				"field1": "value1",
				"field2": "not_an_int",
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

			cc := &ConfigContainer{
				data: tc.data,
			}

			err := cc.Unmarshaler(tc.prefix, tc.obj)
			if (err != nil) != tc.wantErr {
				t.Errorf("Unmarshaler() error = %v, wantErr %v", err, tc.wantErr)
				return
			}
		})
	}
}
