package dartsass

import (
	"fmt"
	"reflect"
	"testing"

	godartsassv1 "github.com/bep/godartsass"
)

func TestLoad_1(t *testing.T) {
	type args struct {
		url string
	}
	tests := []struct {
		name    string
		t       importResolverV1
		args    args
		want    godartsassv1.Import
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

			got, err := tt.t.Load(tt.args.url)
			if (err != nil) != tt.wantErr {
				t.Errorf("importResolverV1.Load() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("importResolverV1.Load() = %v, want %v", got, tt.want)
			}
		})
	}
}
