package mock

import (
	"fmt"
	"testing"

	"github.com/beego/beego/v2/client/orm"
)

func TestValuesList_2(t *testing.T) {
	type args struct {
		results *[]orm.ParamsList
		exprs   []string
	}
	tests := []struct {
		name    string
		d       *DoNothingQuerySetter
		args    args
		want    int64
		wantErr bool
	}{
		{
			name: "test",
			d:    &DoNothingQuerySetter{},
			args: args{
				results: &[]orm.ParamsList{},
				exprs:   []string{},
			},
			want:    0,
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

			got, err := tt.d.ValuesList(tt.args.results, tt.args.exprs...)
			if (err != nil) != tt.wantErr {
				t.Errorf("ValuesList() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("ValuesList() got = %v, want %v", got, tt.want)
			}
		})
	}
}
