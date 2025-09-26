package hugolib

import (
	"fmt"
	"testing"

	"github.com/gohugoio/hugo/langs"
)

func TestIsMultilingual_2(t *testing.T) {
	type args struct {
		h *HugoSites
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "test_case_1",
			args: args{
				h: &HugoSites{
					Sites: []*Site{
						{
							language: &langs.Language{
								Lang: "en",
							},
						},
						{
							language: &langs.Language{
								Lang: "fr",
							},
						},
					},
				},
			},
			want: true,
		},
		{
			name: "test_case_2",
			args: args{
				h: &HugoSites{
					Sites: []*Site{
						{
							language: &langs.Language{
								Lang: "en",
							},
						},
					},
				},
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			if got := tt.args.h.isMultilingual(); got != tt.want {
				t.Errorf("isMultilingual() = %v, want %v", got, tt.want)
			}
		})
	}
}
