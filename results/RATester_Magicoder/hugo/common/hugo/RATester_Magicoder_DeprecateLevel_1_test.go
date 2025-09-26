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
			name: "Test DeprecateLevel with LevelError",
			args: args{
				item:        "item1",
				alternative: "alternative1",
				version:     "version1",
				level:       logg.LevelError,
			},
		},
		{
			name: "Test DeprecateLevel with LevelWarn",
			args: args{
				item:        "item2",
				alternative: "alternative2",
				version:     "version2",
				level:       logg.LevelWarn,
			},
		},
		{
			name: "Test DeprecateLevel with LevelInfo",
			args: args{
				item:        "item3",
				alternative: "alternative3",
				version:     "version3",
				level:       logg.LevelInfo,
			},
		},
		{
			name: "Test DeprecateLevel with LevelDebug",
			args: args{
				item:        "item4",
				alternative: "alternative4",
				version:     "version4",
				level:       logg.LevelDebug,
			},
		},
		{
			name: "Test DeprecateLevel with LevelTrace",
			args: args{
				item:        "item5",
				alternative: "alternative5",
				version:     "version5",
				level:       logg.LevelTrace,
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
