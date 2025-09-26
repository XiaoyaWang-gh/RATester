package page

import (
	"context"
	"fmt"
	"reflect"
	"testing"

	"github.com/gohugoio/hugo/related"
)

func TestGetOrCreateIndex_1(t *testing.T) {
	type args struct {
		ctx context.Context
		p   Pages
	}
	tests := []struct {
		name    string
		s       *RelatedDocsHandler
		args    args
		want    *related.InvertedIndex
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

			got, err := tt.s.getOrCreateIndex(tt.args.ctx, tt.args.p)
			if (err != nil) != tt.wantErr {
				t.Errorf("RelatedDocsHandler.getOrCreateIndex() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("RelatedDocsHandler.getOrCreateIndex() = %v, want %v", got, tt.want)
			}
		})
	}
}
