package bean

import (
	"fmt"
	"reflect"
	"testing"
)

func TestSetUIntXValue_1(t *testing.T) {
	tests := []struct {
		name      string
		dftValue  string
		bitSize   int
		fn        string
		fv        reflect.Value
		wantErr   bool
		wantValue uint64
	}{
		{
			name:      "Test case 1",
			dftValue:  "10",
			bitSize:   64,
			fn:        "testField",
			fv:        reflect.ValueOf(new(uint64)).Elem(),
			wantErr:   false,
			wantValue: 10,
		},
		{
			name:     "Test case 2",
			dftValue: "invalid",
			bitSize:  64,
			fn:       "testField",
			fv:       reflect.ValueOf(new(uint64)).Elem(),
			wantErr:  true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			err := (&TagAutoWireBeanFactory{}).setUIntXValue(tt.dftValue, tt.bitSize, tt.fn, tt.fv)
			if (err != nil) != tt.wantErr {
				t.Errorf("setUIntXValue() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if err != nil {
				return
			}
			if tt.fv.Uint() != tt.wantValue {
				t.Errorf("setUIntXValue() wantValue = %v, got = %v", tt.wantValue, tt.fv.Uint())
			}
		})
	}
}
