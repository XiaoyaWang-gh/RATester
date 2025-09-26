package htime

import (
	"fmt"
	"reflect"
	"testing"
	"time"
)

func TestToTimeInDefaultLocationE_1(t *testing.T) {
	type args struct {
		i        any
		location *time.Location
	}
	tests := []struct {
		name    string
		args    args
		wantTim time.Time
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

			gotTim, err := ToTimeInDefaultLocationE(tt.args.i, tt.args.location)
			if (err != nil) != tt.wantErr {
				t.Errorf("ToTimeInDefaultLocationE() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotTim, tt.wantTim) {
				t.Errorf("ToTimeInDefaultLocationE() = %v, want %v", gotTim, tt.wantTim)
			}
		})
	}
}
