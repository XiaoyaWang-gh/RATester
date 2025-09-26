package media

import (
	"fmt"
	"reflect"
	"testing"
)

func TestFromString_4(t *testing.T) {
	tests := []struct {
		name    string
		input   string
		want    Type
		wantErr bool
	}{
		{
			name:  "valid media type",
			input: "application/json",
			want: Type{
				Type:       "application/json",
				MainType:   "application",
				SubType:    "json",
				mimeSuffix: "",
			},
			wantErr: false,
		},
		{
			name:  "valid media type with suffix",
			input: "application/json+xml",
			want: Type{
				Type:       "application/json+xml",
				MainType:   "application",
				SubType:    "json",
				mimeSuffix: "xml",
			},
			wantErr: false,
		},
		{
			name:    "invalid media type",
			input:   "application",
			want:    Type{},
			wantErr: true,
		},
		{
			name:    "invalid media type with multiple slashes",
			input:   "application/json/xml",
			want:    Type{},
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

			got, err := FromString(tt.input)
			if (err != nil) != tt.wantErr {
				t.Errorf("FromString() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("FromString() = %v, want %v", got, tt.want)
			}
		})
	}
}
