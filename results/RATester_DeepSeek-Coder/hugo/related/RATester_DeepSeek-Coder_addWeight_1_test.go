package related

import (
	"fmt"
	"testing"
)

func TestAddWeight_1(t *testing.T) {
	type fields struct {
		Doc     Document
		Weight  int
		Matches int
	}
	type args struct {
		w int
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   int
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

			r := &rank{
				Doc:     tt.fields.Doc,
				Weight:  tt.fields.Weight,
				Matches: tt.fields.Matches,
			}
			r.addWeight(tt.args.w)
			if got := r.Weight; got != tt.want {
				t.Errorf("addWeight() = %v, want %v", got, tt.want)
			}
		})
	}
}
