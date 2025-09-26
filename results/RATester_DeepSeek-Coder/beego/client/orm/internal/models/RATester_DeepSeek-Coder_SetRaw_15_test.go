package models

import (
	"fmt"
	"testing"
	"time"
)

func TestSetRaw_15(t *testing.T) {
	type args struct {
		value interface{}
	}
	tests := []struct {
		name    string
		e       *DateField
		args    args
		wantErr bool
	}{
		{
			name: "Test with time.Time",
			e:    &DateField{},
			args: args{
				value: time.Now(),
			},
			wantErr: false,
		},
		{
			name: "Test with string",
			e:    &DateField{},
			args: args{
				value: "2022-01-01",
			},
			wantErr: false,
		},
		{
			name: "Test with unknown type",
			e:    &DateField{},
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

			e := tt.e
			if err := e.SetRaw(tt.args.value); (err != nil) != tt.wantErr {
				t.Errorf("DateField.SetRaw() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
