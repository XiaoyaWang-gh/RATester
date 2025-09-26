package minifier

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/gohugoio/hugo/resources"
)

func TestNew_24(t *testing.T) {
	type args struct {
		rs *resources.Spec
	}
	tests := []struct {
		name    string
		args    args
		want    *Client
		wantErr bool
	}{
		{
			name: "TestNew_Success",
			args: args{
				rs: &resources.Spec{},
			},
			want: &Client{
				rs: &resources.Spec{},
			},
			wantErr: false,
		},
		{
			name: "TestNew_Failure",
			args: args{
				rs: nil,
			},
			want:    nil,
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

			got, err := New(tt.args.rs)
			if (err != nil) != tt.wantErr {
				t.Errorf("New() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("New() = %v, want %v", got, tt.want)
			}
		})
	}
}
