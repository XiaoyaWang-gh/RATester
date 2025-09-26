package commands

import (
	"fmt"
	"testing"
)

func TestconvertJekyllContent_1(t *testing.T) {
	type args struct {
		m       any
		content string
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

			c := &importCommand{}
			got, err := c.convertJekyllContent(tt.args.m, tt.args.content)
			if (err != nil) != tt.wantErr {
				t.Errorf("convertJekyllContent() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("convertJekyllContent() = %v, want %v", got, tt.want)
			}
		})
	}
}
