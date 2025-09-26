package pagemeta

import (
	"fmt"
	"reflect"
	"testing"
	"time"
)

func TestDateAndSlugFromBaseFilename_1(t *testing.T) {
	type args struct {
		location *time.Location
		name     string
	}
	tests := []struct {
		name  string
		args  args
		want  time.Time
		want1 string
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

			got, got1 := dateAndSlugFromBaseFilename(tt.args.location, tt.args.name)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("dateAndSlugFromBaseFilename() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("dateAndSlugFromBaseFilename() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}
