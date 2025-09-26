package commands

import (
	"fmt"
	"reflect"
	"testing"
	"time"
)

func TestParseJekyllFilename_1(t *testing.T) {
	type args struct {
		filename string
	}
	tests := []struct {
		name    string
		args    args
		want    time.Time
		want1   string
		wantErr bool
	}{
		{
			name: "Test with valid filename",
			args: args{
				filename: "2022-1-1-test.md",
			},
			want:    time.Date(2022, 1, 1, 0, 0, 0, 0, time.UTC),
			want1:   "test",
			wantErr: false,
		},
		{
			name: "Test with invalid filename",
			args: args{
				filename: "invalid-filename",
			},
			want:    time.Time{},
			want1:   "",
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			c := &importCommand{}
			got, got1, err := c.parseJekyllFilename(tt.args.filename)
			if (err != nil) != tt.wantErr {
				t.Errorf("parseJekyllFilename() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("parseJekyllFilename() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("parseJekyllFilename() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}
