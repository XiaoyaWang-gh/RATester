package hugolib

import (
	"fmt"
	"testing"
)

func TestshiftToOutputFormat_1(t *testing.T) {
	type args struct {
		isRenderingSite bool
		idx             int
	}
	tests := []struct {
		name    string
		args    args
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

			p := &pageState{}
			if err := p.shiftToOutputFormat(tt.args.isRenderingSite, tt.args.idx); (err != nil) != tt.wantErr {
				t.Errorf("shiftToOutputFormat() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
