package metrics

import (
	"fmt"
	"reflect"
	"testing"

	stdprometheus "github.com/prometheus/client_golang/prometheus"
)

func TestToLabels_1(t *testing.T) {
	tests := []struct {
		name string
		lvs  labelNamesValues
		want stdprometheus.Labels
	}{
		{
			name: "empty",
			lvs:  labelNamesValues{},
			want: stdprometheus.Labels{},
		},
		{
			name: "one pair",
			lvs:  labelNamesValues{"name", "value"},
			want: stdprometheus.Labels{"name": "value"},
		},
		{
			name: "multiple pairs",
			lvs:  labelNamesValues{"name1", "value1", "name2", "value2"},
			want: stdprometheus.Labels{"name1": "value1", "name2": "value2"},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			if got := tt.lvs.ToLabels(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ToLabels() = %v, want %v", got, tt.want)
			}
		})
	}
}
