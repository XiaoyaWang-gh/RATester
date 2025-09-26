package paths

import (
	"fmt"
	"testing"

	"github.com/gohugoio/hugo/common/types"
)

func TestNameNoIdentifier_1(t *testing.T) {
	type fields struct {
		s           string
		identifiers []types.LowHigh[string]
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{
			name: "Test with identifiers",
			fields: fields{
				s: "test/path/with/identifiers",
				identifiers: []types.LowHigh[string]{
					{Low: 10, High: 15},
					{Low: 16, High: 20},
				},
			},
			want: "test/path",
		},
		{
			name: "Test without identifiers",
			fields: fields{
				s: "test/path/without/identifiers",
			},
			want: "test/path/without/identifiers",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			p := &Path{
				s:           tt.fields.s,
				identifiers: tt.fields.identifiers,
			}
			if got := p.NameNoIdentifier(); got != tt.want {
				t.Errorf("NameNoIdentifier() = %v, want %v", got, tt.want)
			}
		})
	}
}
