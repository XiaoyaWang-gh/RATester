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
			name: "single",
			lvs:  labelNamesValues{"key1", "value1"},
			want: stdprometheus.Labels{"key1": "value1"},
		},
		{
			name: "multiple",
			lvs:  labelNamesValues{"key1", "value1", "key2", "value2", "key3", "value3"},
			want: stdprometheus.Labels{"key1": "value1", "key2": "value2", "key3": "value3"},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			got := tt.lvs.ToLabels()
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ToLabels() = %v, want %v", got, tt.want)
			}
		})
	}
}
