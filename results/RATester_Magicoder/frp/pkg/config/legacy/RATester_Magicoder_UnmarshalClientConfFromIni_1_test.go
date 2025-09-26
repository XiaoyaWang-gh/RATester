package legacy

import (
	"fmt"
	"reflect"
	"testing"
)

func TestUnmarshalClientConfFromIni_1(t *testing.T) {
	type args struct {
		source interface{}
	}
	tests := []struct {
		name    string
		args    args
		want    ClientCommonConf
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

			got, err := UnmarshalClientConfFromIni(tt.args.source)
			if (err != nil) != tt.wantErr {
				t.Errorf("UnmarshalClientConfFromIni() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("UnmarshalClientConfFromIni() = %v, want %v", got, tt.want)
			}
		})
	}
}
