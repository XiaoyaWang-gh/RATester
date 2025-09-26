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
			name: "acceptNone is true",
			w: Whitelist{
				acceptNone: true,
			},
			want:    []byte(`"acceptNone"`),
			wantErr: false,
		},
		{
			name: "acceptNone is false",
			w: Whitelist{
				acceptNone: false,
				patterns: []*regexp.Regexp{
					regexp.MustCompile("pattern1"),
					regexp.MustCompile("pattern2"),
				},
				patternsStrings: []string{"pattern1", "pattern2"},
			},
			want:    []byte(`["pattern1","pattern2"]`),
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
