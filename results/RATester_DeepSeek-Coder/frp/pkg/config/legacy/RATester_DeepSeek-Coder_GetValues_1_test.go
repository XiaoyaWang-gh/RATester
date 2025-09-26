package legacy

import (
	"fmt"
	"reflect"
	"testing"
)

func TestGetValues_1(t *testing.T) {
	tests := []struct {
		name string
		want *Values
	}{
		{
			name: "Test case 1",
			want: &Values{
				Envs: glbEnvs,
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

			got := GetValues()
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetValues() = %v, want %v", got, tt.want)
			}
		})
	}
}
