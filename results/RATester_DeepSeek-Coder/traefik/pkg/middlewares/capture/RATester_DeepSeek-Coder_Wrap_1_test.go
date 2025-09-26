package capture

import (
	"fmt"
	"net/http"
	"reflect"
	"testing"
)

func TestWrap_1(t *testing.T) {
	type args struct {
		next http.Handler
	}
	tests := []struct {
		name    string
		args    args
		want    http.Handler
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

			got, err := Wrap(tt.args.next)
			if (err != nil) != tt.wantErr {
				t.Errorf("Wrap() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Wrap() = %v, want %v", got, tt.want)
			}
		})
	}
}
