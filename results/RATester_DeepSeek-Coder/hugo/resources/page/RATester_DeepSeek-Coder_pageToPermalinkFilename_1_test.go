package page

import (
	"fmt"
	"testing"
)

func TestPageToPermalinkFilename_1(t *testing.T) {
	type args struct {
		p Page
		a string
	}
	tests := []struct {
		name    string
		l       PermalinkExpander
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

			got, err := tt.l.pageToPermalinkFilename(tt.args.p, tt.args.a)
			if (err != nil) != tt.wantErr {
				t.Errorf("PermalinkExpander.pageToPermalinkFilename() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("PermalinkExpander.pageToPermalinkFilename() = %v, want %v", got, tt.want)
			}
		})
	}
}
