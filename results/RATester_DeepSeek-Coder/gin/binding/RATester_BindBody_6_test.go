package binding

import (
	"fmt"
	"testing"
)

func TestBindBody_6(t *testing.T) {
	b := tomlBinding{}

	type testStruct struct {
		Name string `toml:"name"`
		Age  int    `toml:"age"`
	}

	testCases := []struct {
		name    string
		body    string
		wantErr bool
	}{
		{
			name:    "valid toml",
			body:    "name = \"John\"\nage = 30",
			wantErr: false,
		},
		{
			name:    "invalid toml",
			body:    "name = \"John\"\nage = \"30\"",
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

			obj := &testStruct{}
			err := b.BindBody([]byte(tc.body), obj)
			if (err != nil) != tc.wantErr {
				t.Errorf("BindBody() error = %v, wantErr %v", err, tc.wantErr)
			}
		})
	}
}
