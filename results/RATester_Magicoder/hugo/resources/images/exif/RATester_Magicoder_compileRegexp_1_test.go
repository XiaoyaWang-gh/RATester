package exif

import (
	"fmt"
	"reflect"
	"regexp"
	"testing"
)

func TestcompileRegexp_1(t *testing.T) {
	tests := []struct {
		name    string
		input   string
		want    *regexp.Regexp
		wantErr bool
	}{
		{
			name:  "Test case 1",
			input: "(?i)test",
			want:  regexp.MustCompile("(?i)test"),
		},
		{
			name:  "Test case 2",
			input: "test",
			want:  regexp.MustCompile("(?i)test"),
		},
		{
			name:  "Test case 3",
			input: "",
			want:  nil,
		},
		{
			name:  "Test case 4",
			input: " ",
			want:  nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			got, err := compileRegexp(tt.input)
			if (err != nil) != tt.wantErr {
				t.Errorf("compileRegexp() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("compileRegexp() = %v, want %v", got, tt.want)
			}
		})
	}
}
