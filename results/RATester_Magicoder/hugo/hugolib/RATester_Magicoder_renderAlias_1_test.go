package hugolib

import (
	"fmt"
	"io"
	"reflect"
	"testing"

	"github.com/gohugoio/hugo/resources/page"
)

func TestrenderAlias_1(t *testing.T) {
	type args struct {
		permalink string
		p         page.Page
	}
	tests := []struct {
		name    string
		a       aliasHandler
		args    args
		want    io.Reader
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

			got, err := tt.a.renderAlias(tt.args.permalink, tt.args.p)
			if (err != nil) != tt.wantErr {
				t.Errorf("aliasHandler.renderAlias() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("aliasHandler.renderAlias() = %v, want %v", got, tt.want)
			}
		})
	}
}
