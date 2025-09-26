package langs

import (
	"fmt"
	"testing"
)

func TestloadLocation_1(t *testing.T) {
	type args struct {
		tzStr string
	}
	tests := []struct {
		name    string
		l       *Language
		args    args
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

			if err := tt.l.loadLocation(tt.args.tzStr); (err != nil) != tt.wantErr {
				t.Errorf("Language.loadLocation() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
