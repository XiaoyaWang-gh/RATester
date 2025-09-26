package validation

import (
	"fmt"
	"reflect"
	"testing"
)

func TestGetLimitValue_2(t *testing.T) {
	type fields struct {
		Min Min
		Max Max
		Key string
	}
	tests := []struct {
		name   string
		fields fields
		want   interface{}
	}{
		{
			name: "TestGetLimitValue",
			fields: fields{
				Min: Min{
					Min: 1,
					Key: "Min",
				},
				Max: Max{
					Max: 10,
					Key: "Max",
				},
				Key: "Range",
			},
			want: []int{1, 10},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			r := Range{
				Min: tt.fields.Min,
				Max: tt.fields.Max,
				Key: tt.fields.Key,
			}
			if got := r.GetLimitValue(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetLimitValue() = %v, want %v", got, tt.want)
			}
		})
	}
}
