package page

import (
	"fmt"
	"reflect"
	"testing"
)

func TestMarshalPageToJSON_1(t *testing.T) {
	type args struct {
		p Page
	}
	tests := []struct {
		name    string
		args    args
		want    []byte
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

			got, err := MarshalPageToJSON(tt.args.p)
			if (err != nil) != tt.wantErr {
				t.Errorf("MarshalPageToJSON() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("MarshalPageToJSON() = %v, want %v", got, tt.want)
			}
		})
	}
}
