package middlewares

import (
	"context"
	"fmt"
	"reflect"
	"testing"

	"github.com/traefik/traefik/v3/pkg/config/dynamic"
	"github.com/vulcand/oxy/v2/utils"
)

func TestGetSourceExtractor_1(t *testing.T) {
	type args struct {
		ctx           context.Context
		sourceMatcher *dynamic.SourceCriterion
	}
	tests := []struct {
		name    string
		args    args
		want    utils.SourceExtractor
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

			got, err := GetSourceExtractor(tt.args.ctx, tt.args.sourceMatcher)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetSourceExtractor() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetSourceExtractor() = %v, want %v", got, tt.want)
			}
		})
	}
}
