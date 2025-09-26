package fiber

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/gofiber/fiber/v3/log"
)

func TestlistenConfigDefault_1(t *testing.T) {
	type args struct {
		config []ListenConfig
	}
	tests := []struct {
		name string
		args args
		want ListenConfig
	}{
		{
			name: "Test with empty config",
			args: args{
				config: []ListenConfig{},
			},
			want: ListenConfig{
				ListenerNetwork: NetworkTCP4,
				OnShutdownError: func(err error) {
					log.Fatalf("shutdown: %v", err)
				},
			},
		},
		{
			name: "Test with non-empty config",
			args: args{
				config: []ListenConfig{
					{
						ListenerNetwork: "tcp4",
						OnShutdownError: func(err error) {
							log.Fatalf("shutdown: %v", err)
						},
					},
				},
			},
			want: ListenConfig{
				ListenerNetwork: "tcp4",
				OnShutdownError: func(err error) {
					log.Fatalf("shutdown: %v", err)
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			if got := listenConfigDefault(tt.args.config...); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("listenConfigDefault() = %v, want %v", got, tt.want)
			}
		})
	}
}
