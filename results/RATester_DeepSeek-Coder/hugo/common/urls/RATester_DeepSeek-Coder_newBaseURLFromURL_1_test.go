package urls

import (
	"fmt"
	"net/url"
	"reflect"
	"testing"
)

func TestNewBaseURLFromURL_1(t *testing.T) {
	type args struct {
		u *url.URL
	}
	tests := []struct {
		name    string
		args    args
		want    BaseURL
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

			got, err := newBaseURLFromURL(tt.args.u)
			if (err != nil) != tt.wantErr {
				t.Errorf("newBaseURLFromURL() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("newBaseURLFromURL() = %v, want %v", got, tt.want)
			}
		})
	}
}
