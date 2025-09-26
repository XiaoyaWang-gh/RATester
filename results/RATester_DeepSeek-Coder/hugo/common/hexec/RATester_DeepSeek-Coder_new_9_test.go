package hexec

import (
	"fmt"
	"reflect"
	"testing"
)

func TestNew_9(t *testing.T) {
	type args struct {
		name               string
		fullyQualifiedName string
		arg                []any
	}
	tests := []struct {
		name    string
		e       *Exec
		args    args
		want    Runner
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

			got, err := tt.e.new(tt.args.name, tt.args.fullyQualifiedName, tt.args.arg...)
			if (err != nil) != tt.wantErr {
				t.Errorf("Exec.new() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Exec.new() = %v, want %v", got, tt.want)
			}
		})
	}
}
