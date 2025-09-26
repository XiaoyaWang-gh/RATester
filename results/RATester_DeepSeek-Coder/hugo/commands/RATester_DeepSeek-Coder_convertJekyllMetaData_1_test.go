package commands

import (
	"fmt"
	"reflect"
	"testing"
	"time"
)

func TestConvertJekyllMetaData_1(t *testing.T) {
	type args struct {
		m        any
		postName string
		postDate time.Time
		draft    bool
	}
	tests := []struct {
		name    string
		args    args
		want    any
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

			c := &importCommand{}
			got, err := c.convertJekyllMetaData(tt.args.m, tt.args.postName, tt.args.postDate, tt.args.draft)
			if (err != nil) != tt.wantErr {
				t.Errorf("convertJekyllMetaData() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("convertJekyllMetaData() = %v, want %v", got, tt.want)
			}
		})
	}
}
