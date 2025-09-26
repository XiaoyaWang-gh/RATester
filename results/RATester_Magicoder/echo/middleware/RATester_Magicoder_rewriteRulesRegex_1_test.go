package middleware

import (
	"fmt"
	"reflect"
	"regexp"
	"testing"
)

func TestRewriteRulesRegex_1(t *testing.T) {
	type args struct {
		rewrite map[string]string
	}
	tests := []struct {
		name string
		args args
		want map[*regexp.Regexp]string
	}{
		{
			name: "Test case 1",
			args: args{
				rewrite: map[string]string{
					"*": "test1",
					"^": "test2",
				},
			},
			want: map[*regexp.Regexp]string{
				regexp.MustCompile(".*?$"):    "test1",
				regexp.MustCompile("^test2$"): "test2",
			},
		},
		{
			name: "Test case 2",
			args: args{
				rewrite: map[string]string{
					"test":  "test1",
					"test2": "test2",
				},
			},
			want: map[*regexp.Regexp]string{
				regexp.MustCompile("test$"):  "test1",
				regexp.MustCompile("test2$"): "test2",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			if got := rewriteRulesRegex(tt.args.rewrite); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("rewriteRulesRegex() = %v, want %v", got, tt.want)
			}
		})
	}
}
