package consulcatalog

import (
	"errors"
	"fmt"
	"testing"
)

func TestInit_21(t *testing.T) {
	tests := []struct {
		desc string
		mock func(p *Provider)
		want error
	}{
		{
			desc: "error while parsing default rule",
			mock: func(p *Provider) {
				p.DefaultRule = "invalid rule"
			},
			want: fmt.Errorf("error while parsing default rule: %w", errors.New("template: :1: function \"invalid\" not defined")),
		},
		{
			desc: "success",
			mock: func(p *Provider) {},
			want: nil,
		},
	}

	for _, test := range tests {
		t.Run(test.desc, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			p := &Provider{}
			test.mock(p)

			err := p.Init()

			if err != nil && test.want != nil && err.Error() != test.want.Error() {
				t.Errorf("got %v, want %v", err, test.want)
			}
		})
	}
}
