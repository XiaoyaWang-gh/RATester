package hugo

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/bep/logg"
)

func TestdeprecationLogLevelFromVersion_1(t *testing.T) {
	tests := []struct {
		name string
		ver  string
		want logg.Level
	}{
		{
			name: "test case 1",
			ver:  "1.12",
			want: logg.LevelError,
		},
		{
			name: "test case 2",
			ver:  "1.6",
			want: logg.LevelWarn,
		},
		{
			name: "test case 3",
			ver:  "1.5",
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
