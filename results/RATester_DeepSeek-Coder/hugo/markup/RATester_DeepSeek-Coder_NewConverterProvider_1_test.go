package markup

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/gohugoio/hugo/markup/converter"
)

func TestNewConverterProvider_1(t *testing.T) {
	type args struct {
		cfg converter.ProviderConfig
	}
	tests := []struct {
		name    string
		args    args
		want    ConverterProvider
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

			got, err := NewConverterProvider(tt.args.cfg)
			if (err != nil) != tt.wantErr {
				t.Errorf("NewConverterProvider() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewConverterProvider() = %v, want %v", got, tt.want)
			}
		})
	}
}
