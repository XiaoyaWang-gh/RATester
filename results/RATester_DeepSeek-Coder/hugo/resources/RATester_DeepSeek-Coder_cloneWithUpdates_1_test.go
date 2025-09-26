package resources

import (
	"fmt"
	"reflect"
	"testing"
)

func TestCloneWithUpdates_1(t *testing.T) {
	type args struct {
		u *transformationUpdate
	}
	tests := []struct {
		name    string
		rc      *genericResource
		args    args
		want    baseResource
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

			got, err := tt.rc.cloneWithUpdates(tt.args.u)
			if (err != nil) != tt.wantErr {
				t.Errorf("genericResource.cloneWithUpdates() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("genericResource.cloneWithUpdates() = %v, want %v", got, tt.want)
			}
		})
	}
}
