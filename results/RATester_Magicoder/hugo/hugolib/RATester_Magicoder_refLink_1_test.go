package hugolib

import (
	"fmt"
	"testing"
)

func TestrefLink_1(t *testing.T) {
	type args struct {
		ref          string
		source       any
		relative     bool
		outputFormat string
	}
	tests := []struct {
		name    string
		args    args
		want    string
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

			s := &siteRefLinker{}
			got, err := s.refLink(tt.args.ref, tt.args.source, tt.args.relative, tt.args.outputFormat)
			if (err != nil) != tt.wantErr {
				t.Errorf("siteRefLinker.refLink() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("siteRefLinker.refLink() = %v, want %v", got, tt.want)
			}
		})
	}
}
