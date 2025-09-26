package models

import (
	"fmt"
	"testing"
	"time"

	"github.com/beego/beego/v2/client/orm/internal/utils"
)

func TestSetRaw_12(t *testing.T) {
	type args struct {
		value interface{}
	}
	tests := []struct {
		name    string
		e       *TimeField
		args    args
		wantErr bool
	}{
		{
			name: "Test with time.Time",
			e:    &TimeField{},
			args: args{
				value: time.Now(),
			},
			wantErr: false,
		},
		{
			name: "Test with string",
			e:    &TimeField{},
			args: args{
				value: time.Now().Format(utils.FormatTime),
			},
			wantErr: false,
		},
		{
			name: "Test with unknown value",
			e:    &TimeField{},
			args: args{
				value: 123,
			},
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

			if err := tt.e.SetRaw(tt.args.value); (err != nil) != tt.wantErr {
				t.Errorf("TimeField.SetRaw() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
