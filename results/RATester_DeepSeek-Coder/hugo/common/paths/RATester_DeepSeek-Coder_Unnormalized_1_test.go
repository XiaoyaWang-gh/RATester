package paths

import (
	"fmt"
	"reflect"
	"testing"
)

func TestUnnormalized_1(t *testing.T) {
	type fields struct {
		unnormalized *Path
	}
	tests := []struct {
		name   string
		fields fields
		want   *Path
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			p := &Path{
				unnormalized: tt.fields.unnormalized,
			}
			if got := p.Unnormalized(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Path.Unnormalized() = %v, want %v", got, tt.want)
			}
		})
	}
}
