package hugolib

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/gohugoio/hugo/resources/page"
)

func TestcreateTargetPathDescriptor_1(t *testing.T) {
	type args struct {
		p *pageState
	}
	tests := []struct {
		name    string
		args    args
		want    page.TargetPathDescriptor
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

			got, err := createTargetPathDescriptor(tt.args.p)
			if (err != nil) != tt.wantErr {
				t.Errorf("createTargetPathDescriptor() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("createTargetPathDescriptor() = %v, want %v", got, tt.want)
			}
		})
	}
}
