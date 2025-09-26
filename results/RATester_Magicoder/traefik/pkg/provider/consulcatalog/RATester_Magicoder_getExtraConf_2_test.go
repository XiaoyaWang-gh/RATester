package consulcatalog

import (
	"fmt"
	"reflect"
	"testing"
)

func TestGetExtraConf_2(t *testing.T) {
	type args struct {
		labels map[string]string
	}
	tests := []struct {
		name    string
		p       *Provider
		args    args
		want    configuration
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

			got, err := tt.p.getExtraConf(tt.args.labels)
			if (err != nil) != tt.wantErr {
				t.Errorf("getExtraConf() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("getExtraConf() = %v, want %v", got, tt.want)
			}
		})
	}
}
