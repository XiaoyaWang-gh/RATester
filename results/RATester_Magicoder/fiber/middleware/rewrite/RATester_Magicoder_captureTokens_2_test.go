package rewrite

import (
	"fmt"
	"reflect"
	"regexp"
	"strings"
	"testing"
)

func TestcaptureTokens_2(t *testing.T) {
	type args struct {
		pattern *regexp.Regexp
		input   string
	}
	tests := []struct {
		name string
		args args
		want *strings.Replacer
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			if got := captureTokens(tt.args.pattern, tt.args.input); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("captureTokens() = %v, want %v", got, tt.want)
			}
		})
	}
}
