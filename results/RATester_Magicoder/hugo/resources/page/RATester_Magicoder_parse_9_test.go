package page

import (
	"fmt"
	"reflect"
	"testing"
)

func Testparse_9(t *testing.T) {
	type args struct {
		patterns map[string]string
	}
	tests := []struct {
		name    string
		l       PermalinkExpander
		args    args
		want    map[string]func(Page) (string, error)
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

			got, err := tt.l.parse(tt.args.patterns)
			if (err != nil) != tt.wantErr {
				t.Errorf("PermalinkExpander.parse() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("PermalinkExpander.parse() = %v, want %v", got, tt.want)
			}
		})
	}
}
