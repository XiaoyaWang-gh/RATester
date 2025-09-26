package logs

import (
	"fmt"
	"reflect"
	"testing"
	"time"
)

func TestformatTimeHeader_1(t *testing.T) {
	type args struct {
		when time.Time
	}
	tests := []struct {
		name  string
		args  args
		want  []byte
		want1 int
		want2 int
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

			got, got1, got2 := formatTimeHeader(tt.args.when)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("formatTimeHeader() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("formatTimeHeader() got1 = %v, want %v", got1, tt.want1)
			}
			if got2 != tt.want2 {
				t.Errorf("formatTimeHeader() got2 = %v, want %v", got2, tt.want2)
			}
		})
	}
}
