package param

import (
	"fmt"
	"reflect"
	"testing"
	"time"
)

func Testparse_13(t *testing.T) {
	tp := timeParser{}
	tests := []struct {
		name    string
		value   string
		want    interface{}
		wantErr bool
	}{
		{
			name:    "valid RFC3339",
			value:   "2006-01-02T15:04:05Z",
			want:    time.Date(2006, 1, 2, 15, 4, 5, 0, time.UTC),
			wantErr: false,
		},
		{
			name:    "valid date",
			value:   "2006-01-02",
			want:    time.Date(2006, 1, 2, 0, 0, 0, 0, time.UTC),
			wantErr: false,
		},
		{
			name:    "invalid date",
			value:   "2006-13-02",
			want:    time.Time{},
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

			got, err := tp.parse(tt.value, reflect.TypeOf(time.Time{}))
			if (err != nil) != tt.wantErr {
				t.Errorf("parse() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("parse() = %v, want %v", got, tt.want)
			}
		})
	}
}
