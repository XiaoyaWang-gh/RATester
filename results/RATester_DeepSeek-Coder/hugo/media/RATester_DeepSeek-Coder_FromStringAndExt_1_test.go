package media

import (
	"fmt"
	"reflect"
	"testing"
)

func TestFromStringAndExt_1(t *testing.T) {
	tests := []struct {
		name    string
		input   string
		ext     []string
		want    Type
		wantErr bool
	}{
		{
			name:  "valid media type and extensions",
			input: "application/json",
			ext:   []string{".json", ".js"},
			want: Type{
				Type:        "application/json",
				MainType:    "application",
				SubType:     "json",
				Delimiter:   DefaultDelimiter,
				SuffixesCSV: "json,js",
			},
			wantErr: false,
		},
		{
			name:    "invalid media type",
			input:   "invalid",
			ext:     []string{".json", ".js"},
			want:    Type{},
			wantErr: true,
		},
		{
			name:  "empty extensions",
			input: "application/json",
			ext:   []string{},
			want: Type{
				Type:        "application/json",
				MainType:    "application",
				SubType:     "json",
				Delimiter:   DefaultDelimiter,
				SuffixesCSV: "",
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			got, err := FromStringAndExt(tt.input, tt.ext...)
			if (err != nil) != tt.wantErr {
				t.Errorf("FromStringAndExt() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("FromStringAndExt() = %v, want %v", got, tt.want)
			}
		})
	}
}
