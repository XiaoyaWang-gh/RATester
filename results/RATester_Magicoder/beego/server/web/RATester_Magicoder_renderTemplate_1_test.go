package web

import (
	"bytes"
	"fmt"
	"reflect"
	"testing"
)

func TestrenderTemplate_1(t *testing.T) {
	type args struct {
		c *Controller
	}
	tests := []struct {
		name    string
		args    args
		want    bytes.Buffer
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

			got, err := tt.args.c.renderTemplate()
			if (err != nil) != tt.wantErr {
				t.Errorf("renderTemplate() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("renderTemplate() = %v, want %v", got, tt.want)
			}
		})
	}
}
