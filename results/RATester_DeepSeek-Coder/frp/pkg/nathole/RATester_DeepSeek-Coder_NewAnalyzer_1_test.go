package nathole

import (
	"fmt"
	"reflect"
	"testing"
	"time"
)

func TestNewAnalyzer_1(t *testing.T) {
	type args struct {
		dataReserveDuration time.Duration
	}
	tests := []struct {
		name string
		args args
		want *Analyzer
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

			if got := NewAnalyzer(tt.args.dataReserveDuration); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewAnalyzer() = %v, want %v", got, tt.want)
			}
		})
	}
}
