package hugo

import (
	"fmt"
	"testing"

	"github.com/bep/logg"
)

func TestDeprecateLevel_1(t *testing.T) {
	type args struct {
		item        string
		alternative string
		version     string
		level       logg.Level
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "TestDeprecateLevel",
			args: args{
				item:        "item",
				alternative: "alternative",
				version:     "version",
				level:       logg.LevelError,
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

			DeprecateLevel(tt.args.item, tt.args.alternative, tt.args.version, tt.args.level)
		})
	}
}
