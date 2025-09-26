package filecache

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/gohugoio/hugo/helpers"
)

func TestNewCaches_1(t *testing.T) {
	type args struct {
		p *helpers.PathSpec
	}
	tests := []struct {
		name    string
		args    args
		want    Caches
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

			got, err := NewCaches(tt.args.p)
			if (err != nil) != tt.wantErr {
				t.Errorf("NewCaches() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewCaches() = %v, want %v", got, tt.want)
			}
		})
	}
}
