package context

import (
	"fmt"
	"testing"
)

func TestBindJSON_1(t *testing.T) {
	type testStruct struct {
		Name string `json:"name"`
		Age  int    `json:"age"`
	}

	testCases := []struct {
		name    string
		input   *BeegoInput
		wantErr bool
	}{
		{
			name: "valid JSON",
			input: &BeegoInput{
				RequestBody: []byte(`{"name": "John", "age": 30}`),
			},
			wantErr: false,
		},
		{
			name: "invalid JSON",
			input: &BeegoInput{
				RequestBody: []byte(`{"name": "John", "age": "thirty"}`),
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

			ctx := &Context{
				Input: tc.input,
			}

			var got testStruct
			err := ctx.BindJSON(&got)

			if (err != nil) != tc.wantErr {
				t.Errorf("BindJSON() error = %v, wantErr %v", err, tc.wantErr)
			}

			if !tc.wantErr && got.Name != "John" {
				t.Errorf("BindJSON() got = %v, want %v", got, "John")
			}

			if !tc.wantErr && got.Age != 30 {
				t.Errorf("BindJSON() got = %v, want %v", got, 30)
			}
		})
	}
}
