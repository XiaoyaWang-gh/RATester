package legacy

import (
	"fmt"
	"reflect"
	"testing"
)

func TestUnmarshalServerConfFromIni_1(t *testing.T) {
	type args struct {
		source interface{}
	}
	tests := []struct {
		name    string
		args    args
		want    ServerCommonConf
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

			got, err := UnmarshalServerConfFromIni(tt.args.source)
			if (err != nil) != tt.wantErr {
				t.Errorf("UnmarshalServerConfFromIni() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("UnmarshalServerConfFromIni() = %v, want %v", got, tt.want)
			}
		})
	}
}
