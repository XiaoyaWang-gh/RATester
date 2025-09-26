package pagesfromdata

import (
	"fmt"
	"reflect"
	"testing"
)

func TesttoPathMap_1(t *testing.T) {
	ctx := &pagesFromDataTemplateContext{
		p: &PagesFromTemplate{},
	}

	type testCase struct {
		name    string
		input   any
		want    string
		wantMap map[string]any
		wantErr bool
	}

	testCases := []testCase{
		{
			name:  "valid input",
			input: map[string]any{"path": "valid-path"},
			want:  "valid-path",
			wantMap: map[string]any{
				"path": "valid-path",
			},
			wantErr: false,
		},
		{
			name:    "invalid input",
			input:   map[string]any{"path": ""},
			want:    "",
			wantMap: nil,
			wantErr: true,
		},
		{
			name:    "missing path",
			input:   map[string]any{},
			want:    "",
			wantMap: nil,
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

			got, gotMap, err := ctx.toPathMap(tc.input)
			if (err != nil) != tc.wantErr {
				t.Errorf("toPathMap() error = %v, wantErr %v", err, tc.wantErr)
				return
			}
			if got != tc.want {
				t.Errorf("toPathMap() got = %v, want %v", got, tc.want)
			}
			if !reflect.DeepEqual(gotMap, tc.wantMap) {
				t.Errorf("toPathMap() gotMap = %v, want %v", gotMap, tc.wantMap)
			}
		})
	}
}
