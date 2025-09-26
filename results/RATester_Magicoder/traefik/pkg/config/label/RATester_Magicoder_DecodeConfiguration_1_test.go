package label

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/traefik/genconf/dynamic"
)

func TestDecodeConfiguration_1(t *testing.T) {
	type args struct {
		labels map[string]string
	}
	tests := []struct {
		name    string
		args    args
		want    *dynamic.Configuration
		wantErr bool
	}{
		{
			name: "Test case 1",
			args: args{
				labels: map[string]string{
					"traefik.http.routers.my-router.rule": "Host(`example.com`)",
				},
			},
			want: &dynamic.Configuration{
				HTTP: &dynamic.HTTPConfiguration{
					Routers: map[string]*dynamic.Router{
						"my-router": {
							Rule: "Host(`example.com`)",
						},
					},
				},
			},
			wantErr: false,
		},
		// Add more test cases as needed
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			got, err := DecodeConfiguration(tt.args.labels)
			if (err != nil) != tt.wantErr {
				t.Errorf("DecodeConfiguration() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("DecodeConfiguration() = %v, want %v", got, tt.want)
			}
		})
	}
}
