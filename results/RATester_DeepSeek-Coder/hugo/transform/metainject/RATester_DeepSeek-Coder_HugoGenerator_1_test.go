package metainject

import (
	"fmt"
	"testing"

	"github.com/gohugoio/hugo/transform"
)

func TestHugoGenerator_1(t *testing.T) {
	type args struct {
		ft transform.FromTo
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

			if err := HugoGenerator(tt.args.ft); (err != nil) != tt.wantErr {
				t.Errorf("HugoGenerator() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
