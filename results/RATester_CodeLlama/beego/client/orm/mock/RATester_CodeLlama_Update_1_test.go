package mock

import (
	"fmt"
	"testing"

	"github.com/beego/beego/v2/client/orm"
)

func TestUpdate_1(t *testing.T) {
	type args struct {
		values orm.Params
	}
	tests := []struct {
		name    string
		d       *DoNothingQuerySetter
		args    args
		want    int64
		wantErr bool
	}{
		{
			name: "test1",
			d:    &DoNothingQuerySetter{},
			args: args{
				values: orm.Params{
					"name": "test",
				},
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

			got, err := tt.d.Update(tt.args.values)
			if (err != nil) != tt.wantErr {
				t.Errorf("DoNothingQuerySetter.Update() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("DoNothingQuerySetter.Update() = %v, want %v", got, tt.want)
			}
		})
	}
}
