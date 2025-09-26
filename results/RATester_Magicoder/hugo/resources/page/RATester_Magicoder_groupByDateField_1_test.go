package page

import (
	"fmt"
	"reflect"
	"testing"
	"time"
)

func TestgroupByDateField_1(t *testing.T) {
	type args struct {
		format  string
		sorter  func(p Pages) Pages
		getDate func(p Page) time.Time
		order   []string
	}
	tests := []struct {
		name    string
		p       Pages
		args    args
		want    PagesGroup
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

			got, err := tt.p.groupByDateField(tt.args.format, tt.args.sorter, tt.args.getDate, tt.args.order...)
			if (err != nil) != tt.wantErr {
				t.Errorf("groupByDateField() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("groupByDateField() = %v, want %v", got, tt.want)
			}
		})
	}
}
