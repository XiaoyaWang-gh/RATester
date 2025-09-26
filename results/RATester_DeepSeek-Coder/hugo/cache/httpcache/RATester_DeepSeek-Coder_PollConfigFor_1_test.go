package httpcache

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/gohugoio/hugo/common/predicate"
)

func TestPollConfigFor_1(t *testing.T) {
	tests := []struct {
		name     string
		compiled *ConfigCompiled
		arg      string
		want     PollConfigCompiled
	}{
		{
			name: "should return PollConfigCompiled when it exists",
			compiled: &ConfigCompiled{
				PollConfigs: []PollConfigCompiled{
					{
						For: predicate.P[string](func(s string) bool {
							return s == "test"
						}),
					},
				},
			},
			arg: "test",
			want: PollConfigCompiled{
				For: predicate.P[string](func(s string) bool {
					return s == "test"
				}),
			},
		},
		{
			name: "should return empty PollConfigCompiled when it does not exist",
			compiled: &ConfigCompiled{
				PollConfigs: []PollConfigCompiled{
					{
						For: predicate.P[string](func(s string) bool {
							return s == "test"
						}),
					},
				},
			},
			arg:  "not-exist",
			want: PollConfigCompiled{},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			got := tt.compiled.PollConfigFor(tt.arg)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("got %v, want %v", got, tt.want)
			}
		})
	}
}
