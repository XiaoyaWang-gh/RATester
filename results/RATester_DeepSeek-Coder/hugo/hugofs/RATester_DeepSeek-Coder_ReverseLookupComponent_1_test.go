package hugofs

import (
	"fmt"
	"reflect"
	"testing"
)

func TestReverseLookupComponent_1(t *testing.T) {
	type args struct {
		component string
		filename  string
	}
	tests := []struct {
		name    string
		fs      *RootMappingFs
		args    args
		want    []ComponentPath
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

			got, err := tt.fs.ReverseLookupComponent(tt.args.component, tt.args.filename)
			if (err != nil) != tt.wantErr {
				t.Errorf("RootMappingFs.ReverseLookupComponent() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("RootMappingFs.ReverseLookupComponent() = %v, want %v", got, tt.want)
			}
		})
	}
}
