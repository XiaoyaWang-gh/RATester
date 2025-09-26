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
			name: "Test Case 1",
			fields: fields{
				Min: Min{
					Min: 1,
					Key: "min",
				},
				Max: Max{
					Max: 10,
					Key: "max",
				},
				Key: "range",
			},
			want: []int{1, 10},
		},
		// Add more test cases here
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
