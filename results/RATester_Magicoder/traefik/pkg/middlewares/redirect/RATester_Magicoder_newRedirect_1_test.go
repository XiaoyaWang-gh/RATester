package redirect

import (
	"fmt"
	"net/http"
	"reflect"
	"testing"
)

func TestNewRedirect_1(t *testing.T) {
	type args struct {
		next        http.Handler
		regex       string
		replacement string
		permanent   bool
		rawURL      func(*http.Request) string
		name        string
	}
	tests := []struct {
		name    string
		args    args
		want    *redirect
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

			got, err := newRedirect(tt.args.next, tt.args.regex, tt.args.replacement, tt.args.permanent, tt.args.rawURL, tt.args.name)
			if (err != nil) != tt.wantErr {
				t.Errorf("newRedirect() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("newRedirect() = %v, want %v", got, tt.want)
			}
		})
	}
}
