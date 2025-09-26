package paths

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/gohugoio/hugo/common/types"
)

func TestIdentifiers_1(t *testing.T) {
	type fields struct {
		s           string
		identifiers []types.LowHigh[string]
	}
	tests := []struct {
		name   string
		fields fields
		want   []string
	}{
		{
			name: "Test case 1",
			fields: fields{
				s: "test",
				identifiers: []types.LowHigh[string]{
					{Low: 0, High: 1},
					{Low: 1, High: 2},
				},
			},
			want: []string{"t", "es"},
		},
		{
			name: "Test case 2",
			fields: fields{
				s: "test",
				identifiers: []types.LowHigh[string]{
					{Low: 0, High: 2},
					{Low: 2, High: 4},
				},
			},
			want: []string{"te", "st"},
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
			if got := p.Identifiers(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Identifiers() = %v, want %v", got, tt.want)
			}
		})
	}
}
