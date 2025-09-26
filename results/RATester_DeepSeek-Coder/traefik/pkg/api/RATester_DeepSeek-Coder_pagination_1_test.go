package api

import (
	"fmt"
	"net/http"
	"reflect"
	"testing"
)

func TestPagination_1(t *testing.T) {
	type args struct {
		request *http.Request
		maximum int
	}
	tests := []struct {
		name    string
		args    args
		want    pageInfo
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

			got, err := pagination(tt.args.request, tt.args.maximum)
			if (err != nil) != tt.wantErr {
				t.Errorf("pagination() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("pagination() = %v, want %v", got, tt.want)
			}
		})
	}
}
