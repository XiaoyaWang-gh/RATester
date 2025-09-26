package page

import (
	"fmt"
	"testing"
)

func TestpageToPermalinkDate_1(t *testing.T) {
	type args struct {
		p         Page
		dateField string
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

			got, err := tt.l.pageToPermalinkDate(tt.args.p, tt.args.dateField)
			if (err != nil) != tt.wantErr {
				t.Errorf("PermalinkExpander.pageToPermalinkDate() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("PermalinkExpander.pageToPermalinkDate() = %v, want %v", got, tt.want)
			}
		})
	}
}
