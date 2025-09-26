package goldmark

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/gohugoio/hugo/markup/converter"
)

func TestNew_28(t *testing.T) {
	type args struct {
		cfg converter.ProviderConfig
	}
	tests := []struct {
		name    string
		p       provide
		args    args
		want    converter.Provider
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

			got, err := tt.p.New(tt.args.cfg)
			if (err != nil) != tt.wantErr {
				t.Errorf("provide.New() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("provide.New() = %v, want %v", got, tt.want)
			}
		})
	}
}
