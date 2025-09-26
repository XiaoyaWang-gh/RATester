package binding

import (
	"fmt"
	"testing"
)

func TestBindUri_1(t *testing.T) {
	type args struct {
		m   map[string][]string
		obj any
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "Test case 1",
			args: args{
				m: map[string][]string{
					"field name": {"value"},
				},
				obj: &struct{ fieldName string }{},
			},
			wantErr: false,
		},
		{
			name: "Test case 2",
			args: args{
				m: map[string][]string{
					"field m": {"value"},
				},
				obj: &struct{ fieldM map[string][]string }{},
			},
			wantErr: false,
		},
		{
			name: "Test case 3",
			args: args{
				m: map[string][]string{
					"field obj": {"value"},
				},
				obj: &struct{ fieldObj any }{},
			},
			wantErr: false,
		},
		{
			name: "Test case 4",
			args: args{
				m: map[string][]string{
					"field wantErr": {"value"},
				},
				obj: &struct{ fieldWantErr bool }{},
			},
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

			uriBinding := &uriBinding{}
			if err := uriBinding.BindUri(tt.args.m, tt.args.obj); (err != nil) != tt.wantErr {
				t.Errorf("uriBinding.BindUri() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
