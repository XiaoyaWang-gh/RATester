package binder

import (
	"fmt"
	"testing"
)

func TestParse_1(t *testing.T) {
	type testStruct struct {
		Name string `json:"name"`
		Age  int    `json:"age"`
	}

	type testMap map[string]any

	tests := []struct {
		name    string
		alias   string
		out     any
		data    map[string][]string
		wantErr bool
	}{
		{
			name:  "parse to struct",
			alias: "json",
			out:   &testStruct{},
			data:  map[string][]string{"name": {"John"}, "age": {"30"}},
		},
		{
			name:  "parse to map",
			alias: "json",
			out:   &testMap{},
			data:  map[string][]string{"name": {"John"}, "age": {"30"}},
		},
		{
			name:    "invalid out type",
			alias:   "json",
			out:     123,
			data:    map[string][]string{"name": {"John"}, "age": {"30"}},
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

			err := parse(tt.alias, tt.out, tt.data)
			if (err != nil) != tt.wantErr {
				t.Errorf("parse() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
