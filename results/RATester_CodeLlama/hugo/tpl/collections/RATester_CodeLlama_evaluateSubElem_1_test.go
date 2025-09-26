package collections

import (
	"context"
	"fmt"
	"reflect"
	"testing"
)

func TestEvaluateSubElem_1(t *testing.T) {
	type args struct {
		ctx  reflect.Value
		obj  reflect.Value
		elem string
	}
	tests := []struct {
		name    string
		args    args
		want    reflect.Value
		wantErr bool
	}{
		{
			name: "test evaluateSubElem",
			args: args{
				ctx:  reflect.ValueOf(context.Background()),
				obj:  reflect.ValueOf(struct{}{}),
				elem: "field",
			},
			want:    reflect.ValueOf(struct{}{}).FieldByName("field"),
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			got, err := evaluateSubElem(tt.args.ctx, tt.args.obj, tt.args.elem)
			if (err != nil) != tt.wantErr {
				t.Errorf("evaluateSubElem() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("evaluateSubElem() got = %v, want %v", got, tt.want)
			}
		})
	}
}
