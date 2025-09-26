package collections

import (
	"fmt"
	"reflect"
	"testing"
)

func TestevaluateSubElem_1(t *testing.T) {
	type args struct {
		ctx      reflect.Value
		obj      reflect.Value
		elemName string
	}
	tests := []struct {
		name    string
		args    args
		want    reflect.Value
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

			got, err := evaluateSubElem(tt.args.ctx, tt.args.obj, tt.args.elemName)
			if (err != nil) != tt.wantErr {
				t.Errorf("evaluateSubElem() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("evaluateSubElem() = %v, want %v", got, tt.want)
			}
		})
	}
}
