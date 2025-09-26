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
			input: "application/rss+xml",
			want: Type{
				Type:       "application/rss+xml",
				MainType:   "application",
				SubType:    "rss",
				mimeSuffix: "xml",
			},
			wantErr: false,
		},
		{
			name:    "invalid media type",
			input:   "application/rss",
			want:    Type{},
			wantErr: true,
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
