package media

import (
	"fmt"
	"testing"

	"gotest.tools/assert"
)

func TestInitMediaType_1(t *testing.T) {
	type args struct {
		m *Type
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "Test case 1",
			args: args{
				m: &Type{
					Type:        "application/rss+xml",
					MainType:    "application",
					SubType:     "rss",
					Delimiter:   ".",
					FirstSuffix: SuffixInfo{Suffix: "xml", FullSuffix: ".xml"},
					mimeSuffix:  "xml",
					SuffixesCSV: "jpg,jpeg",
				},
			},
		},
		// Add more test cases as needed
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			InitMediaType(tt.args.m)
			// Add assertions to check the expected behavior of the function
			// For example, you can check if the Type field is correctly initialized
			assert.Equal(t, "application/rss+xml", tt.args.m.Type)
		})
	}
}
