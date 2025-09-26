package deployconfig

import (
	"fmt"
	"reflect"
	"testing"
)

func TestParseIncludeExclude_1(t *testing.T) {
	tests := []struct {
		name string
		tgt  *Target
		want error
	}{
		{
			name: "include",
			tgt: &Target{
				Include: "*.html",
			},
			want: nil,
		},
		{
			name: "exclude",
			tgt: &Target{
				Exclude: "*.html",
			},
			want: nil,
		},
		{
			name: "include and exclude",
			tgt: &Target{
				Include: "*.html",
				Exclude: "*.html",
			},
			want: nil,
		},
		{
			name: "invalid include",
			tgt: &Target{
				Include: "**",
			},
			want: fmt.Errorf("invalid deployment.target.include %q: invalid pattern: **", "**"),
		},
		{
			name: "invalid exclude",
			tgt: &Target{
				Exclude: "**",
			},
			want: fmt.Errorf("invalid deployment.target.exclude %q: invalid pattern: **", "**"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			if err := tt.tgt.ParseIncludeExclude(); !reflect.DeepEqual(err, tt.want) {
				t.Errorf("ParseIncludeExclude() error = %v, want %v", err, tt.want)
			}
		})
	}
}
