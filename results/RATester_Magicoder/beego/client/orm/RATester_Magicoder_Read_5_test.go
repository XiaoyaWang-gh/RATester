package orm

import (
	"fmt"
	"testing"

	"github.com/huaweicloud/huaweicloud-sdk-go-v3/services/cce/v3/model"
)

func TestRead_5(t *testing.T) {
	type args struct {
		md   interface{}
		cols []string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "Test case 1",
			args: args{
				md:   &model.User{},
				cols: []string{"id", "name"},
			},
			wantErr: false,
		},
		{
			name: "Test case 2",
			args: args{
				md:   &model.User{},
				cols: []string{"id", "email"},
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			f := &filterOrmDecorator{}
			if err := f.Read(tt.args.md, tt.args.cols...); (err != nil) != tt.wantErr {
				t.Errorf("filterOrmDecorator.Read() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
