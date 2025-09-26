package binding

import (
	"fmt"
	"testing"
)

func TestMapURI_1(t *testing.T) {
	type args struct {
		ptr any
		m   map[string][]string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "test1",
			args: args{
				ptr: &struct {
					str string
				}{},
				m: map[string][]string{
					"str": {"test"},
				},
			},
			wantErr: false,
		},
		{
			name: "test2",
			args: args{
				ptr: &struct {
					str string
				}{},
				m: map[string][]string{
					"str": {"test"},
				},
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

			if err := mapURI(tt.args.ptr, tt.args.m); (err != nil) != tt.wantErr {
				t.Errorf("mapURI() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
