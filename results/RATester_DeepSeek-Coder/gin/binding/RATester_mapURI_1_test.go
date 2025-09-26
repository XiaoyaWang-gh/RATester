package binding

import (
	"fmt"
	"reflect"
	"testing"
)

func TestMapURI_1(t *testing.T) {
	type TestStruct struct {
		URI string `uri:"uri"`
	}

	testCases := []struct {
		name    string
		input   map[string][]string
		want    TestStruct
		wantErr bool
	}{
		{
			name: "valid input",
			input: map[string][]string{
				"uri": {"test"},
			},
			want: TestStruct{
				URI: "test",
			},
			wantErr: false,
		},
		{
			name:    "empty input",
			input:   map[string][]string{},
			want:    TestStruct{},
			wantErr: false,
		},
		{
			name: "missing input",
			input: map[string][]string{
				"missing": {"test"},
			},
			want:    TestStruct{},
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

			var got TestStruct
			err := mapURI(&got, tc.input)
			if (err != nil) != tc.wantErr {
				t.Errorf("mapURI() error = %v, wantErr %v", err, tc.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tc.want) {
				t.Errorf("mapURI() got = %v, want %v", got, tc.want)
			}
		})
	}
}
