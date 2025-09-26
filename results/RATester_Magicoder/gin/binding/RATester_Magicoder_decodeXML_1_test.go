package binding

import (
	"fmt"
	"strings"
	"testing"
)

func TestdecodeXML_1(t *testing.T) {
	type testCase struct {
		name    string
		input   string
		wantErr bool
	}

	testCases := []testCase{
		{
			name:    "valid XML",
			input:   `<root><field name="value"></field></root>`,
			wantErr: false,
		},
		{
			name:    "invalid XML",
			input:   `<root><field name="value"></field>`,
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

			r := strings.NewReader(tc.input)
			var obj any
			err := decodeXML(r, &obj)
			if (err != nil) != tc.wantErr {
				t.Errorf("decodeXML() error = %v, wantErr %v", err, tc.wantErr)
				return
			}
		})
	}
}
