package hugo

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/bep/logg"
)

func TestDeprecationLogLevelFromVersion_1(t *testing.T) {
	tests := []struct {
		name string
		ver  string
		want logg.Level
	}{
		{
			name: "12 minor versions away",
			ver:  "1.0.0",
			want: logg.LevelError,
		},
		{
			name: "6 minor versions away",
			ver:  "1.0.0",
			want: logg.LevelWarn,
		},
		{
			name: "less than 6 minor versions away",
			ver:  "1.0.0",
			want: logg.LevelInfo,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			if got := deprecationLogLevelFromVersion(tt.ver); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("deprecationLogLevelFromVersion() = %v, want %v", got, tt.want)
			}
		})
	}
}
