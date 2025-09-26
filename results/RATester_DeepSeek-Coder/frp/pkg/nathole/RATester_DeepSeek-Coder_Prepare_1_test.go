package nathole

import (
	"fmt"
	"reflect"
	"testing"
)

func TestPrepare_1(t *testing.T) {
	type args struct {
		stunServers []string
	}
	tests := []struct {
		name    string
		args    args
		want    *PrepareResult
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

			got, err := Prepare(tt.args.stunServers)
			if (err != nil) != tt.wantErr {
				t.Errorf("Prepare() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Prepare() = %v, want %v", got, tt.want)
			}
		})
	}
}
