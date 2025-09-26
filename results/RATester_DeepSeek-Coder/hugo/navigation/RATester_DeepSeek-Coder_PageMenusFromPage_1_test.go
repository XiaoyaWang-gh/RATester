package navigation

import (
	"fmt"
	"reflect"
	"testing"
)

func TestPageMenusFromPage_1(t *testing.T) {
	type args struct {
		ms any
		p  Page
	}
	tests := []struct {
		name    string
		args    args
		want    PageMenus
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

			got, err := PageMenusFromPage(tt.args.ms, tt.args.p)
			if (err != nil) != tt.wantErr {
				t.Errorf("PageMenusFromPage() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("PageMenusFromPage() = %v, want %v", got, tt.want)
			}
		})
	}
}
