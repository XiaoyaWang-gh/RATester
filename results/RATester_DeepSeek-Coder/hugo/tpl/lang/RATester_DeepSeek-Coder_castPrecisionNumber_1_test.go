package lang

import (
	"fmt"
	"testing"
)

func TestCastPrecisionNumber_1(t *testing.T) {
	type args struct {
		precision any
		number    any
	}
	tests := []struct {
		name    string
		args    args
		want    uint64
		want1   float64
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

			ns := &Namespace{}
			got, got1, err := ns.castPrecisionNumber(tt.args.precision, tt.args.number)
			if (err != nil) != tt.wantErr {
				t.Errorf("Namespace.castPrecisionNumber() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Namespace.castPrecisionNumber() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("Namespace.castPrecisionNumber() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}
