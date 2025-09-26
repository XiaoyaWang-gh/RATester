package logs

import (
	"fmt"
	"reflect"
	"testing"
)

func TestNewSLACKWriter_1(t *testing.T) {
	type args struct {
		config string
	}
	tests := []struct {
		name string
		args args
		want Logger
	}{
		{
			name: "test1",
			args: args{
				config: "test1",
			},
			want: &SLACKWriter{
				WebhookURL: "test1",
				Level:      LevelTrace,
				formatter:  &SLACKWriter{},
				Formatter:  "test1",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			if got := newSLACKWriter(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("newSLACKWriter() = %v, want %v", got, tt.want)
			}
		})
	}
}
