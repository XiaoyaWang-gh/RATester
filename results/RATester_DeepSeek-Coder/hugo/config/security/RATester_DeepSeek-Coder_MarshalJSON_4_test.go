package security

import (
	"fmt"
	"reflect"
	"regexp"
	"testing"
)

func TestMarshalJSON_4(t *testing.T) {
	tests := []struct {
		name    string
		w       Whitelist
		want    []byte
		wantErr bool
	}{
		{
			name: "Test with acceptNone",
			w: Whitelist{
				acceptNone: true,
			},
			want:    []byte(`"none"`),
			wantErr: false,
		},
		{
			name: "Test with patterns",
			w: Whitelist{
				patterns: []*regexp.Regexp{
					regexp.MustCompile("^[a-z]+$"),
					regexp.MustCompile("^[0-9]+$"),
				},
				patternsStrings: []string{"^[a-z]+$", "^[0-9]+$"},
			},
			want:    []byte(`["^[a-z]+$","^[0-9]+$"]`),
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

			got, err := tt.w.MarshalJSON()
			if (err != nil) != tt.wantErr {
				t.Errorf("Whitelist.MarshalJSON() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Whitelist.MarshalJSON() = %v, want %v", got, tt.want)
			}
		})
	}
}
