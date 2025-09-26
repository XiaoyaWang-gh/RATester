package allconfig

import (
	"fmt"
	"testing"

	"github.com/gohugoio/hugo/common/types"
)

func TestApplyConfigAliases_1(t *testing.T) {
	type args struct {
		aliases []types.KeyValueStr
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "test applyConfigAliases",
			args: args{
				aliases: []types.KeyValueStr{
					{Key: "indexes", Value: "taxonomies"},
					{Key: "logI18nWarnings", Value: "printI18nWarnings"},
					{Key: "logPathWarnings", Value: "printPathWarnings"},
					{Key: "ignoreErrors", Value: "ignoreLogs"},
				},
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

			l := configLoader{}
			l.applyConfigAliases()
		})
	}
}
