package hugolib

import (
	"fmt"
	"testing"

	"github.com/gohugoio/hugo/common/predicate"
)

func TestforEachPage_1(t *testing.T) {
	type args struct {
		include predicate.P[*pageState]
		fn      func(p *pageState) (bool, error)
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
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

			m := &pageMap{}
			if err := m.forEachPage(tt.args.include, tt.args.fn); (err != nil) != tt.wantErr {
				t.Errorf("forEachPage() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
