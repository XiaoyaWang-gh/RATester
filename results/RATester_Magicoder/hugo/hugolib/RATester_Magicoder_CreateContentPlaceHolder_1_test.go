package hugolib

import (
	"fmt"
	"testing"
)

func TestCreateContentPlaceHolder_1(t *testing.T) {
	type args struct {
		filename string
		force    bool
	}
	tests := []struct {
		name    string
		f       ContentFactory
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

			got, err := tt.f.CreateContentPlaceHolder(tt.args.filename, tt.args.force)
			if (err != nil) != tt.wantErr {
				t.Errorf("ContentFactory.CreateContentPlaceHolder() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("ContentFactory.CreateContentPlaceHolder() = %v, want %v", got, tt.want)
			}
		})
	}
}
