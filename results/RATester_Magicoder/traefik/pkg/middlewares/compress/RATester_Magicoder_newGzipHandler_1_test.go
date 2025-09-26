package compress

import (
	"fmt"
	"net/http"
	"reflect"
	"testing"
)

func TestNewGzipHandler_1(t *testing.T) {
	type args struct {
		c *compress
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

			got, err := tt.args.c.newGzipHandler()
			if (err != nil) != tt.wantErr {
				t.Errorf("newGzipHandler() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("newGzipHandler() = %v, want %v", got, tt.want)
			}
		})
	}
}
