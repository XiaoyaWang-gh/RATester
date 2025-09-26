package htesting

import (
	"fmt"
	"reflect"
	"regexp"
	"testing"

	"github.com/frankban/quicktest"
)

func TestNewPinnedRunner_1(t *testing.T) {
	type args struct {
		t            testing.TB
		pinnedTestRe string
	}
	tests := []struct {
		name string
		args args
		want *PinnedRunner
	}{
		{
			name: "Test case 1",
			args: args{
				t:            &testing.T{},
				pinnedTestRe: ".*",
			},
			want: &PinnedRunner{
				c:  &quicktest.C{},
				re: regexp.MustCompile("(?i).*"),
			},
		},
		{
			name: "Test case 2",
			args: args{
				t:            &testing.T{},
				pinnedTestRe: "test",
			},
			want: &PinnedRunner{
				c:  &quicktest.C{},
				re: regexp.MustCompile("(?i)test"),
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

			if got := NewPinnedRunner(tt.args.t, tt.args.pinnedTestRe); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewPinnedRunner() = %v, want %v", got, tt.want)
			}
		})
	}
}
