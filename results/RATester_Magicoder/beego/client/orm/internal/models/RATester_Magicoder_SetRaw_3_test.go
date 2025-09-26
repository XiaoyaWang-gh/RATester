package models

import (
	"fmt"
	"testing"
	"time"
)

func TestSetRaw_3(t *testing.T) {
	type args struct {
		value interface{}
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "Test with time.Time",
			args: args{
				value: time.Now(),
			},
			wantErr: false,
		},
		{
			name: "Test with string",
			args: args{
				value: "2006-01-02 15:04:05",
			},
			wantErr: false,
		},
		{
			name: "Test with unknown type",
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

			e := &DateTimeField{}
			if err := e.SetRaw(tt.args.value); (err != nil) != tt.wantErr {
				t.Errorf("DateTimeField.SetRaw() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
