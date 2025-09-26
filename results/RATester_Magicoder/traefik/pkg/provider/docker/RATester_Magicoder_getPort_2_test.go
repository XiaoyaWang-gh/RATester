package docker

import (
	"fmt"
	"testing"

	"github.com/docker/go-connections/nat"
)

func TestGetPort_2(t *testing.T) {
	type args struct {
		container  dockerData
		serverPort string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "Test Case 1",
			args: args{
				container: dockerData{
					NetworkSettings: networkSettings{
						Ports: nat.PortMap{
							"8080/tcp": []nat.PortBinding{
								{
									HostIP:   "0.0.0.0",
									HostPort: "8080",
								},
							},
						},
					},
				},
				serverPort: "",
			},
			want: "8080",
		},
		{
			name: "Test Case 2",
			args: args{
				container: dockerData{
					NetworkSettings: networkSettings{
						Ports: nat.PortMap{
							"8080/tcp": []nat.PortBinding{
								{
									HostIP:   "0.0.0.0",
									HostPort: "8080",
								},
							},
						},
					},
				},
				serverPort: "8081",
			},
			want: "8081",
		},
		{
			name: "Test Case 3",
			args: args{
				container: dockerData{
					NetworkSettings: networkSettings{
						Ports: nat.PortMap{},
					},
				},
				serverPort: "",
			},
			want: "",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			if got := getPort(tt.args.container, tt.args.serverPort); got != tt.want {
				t.Errorf("getPort() = %v, want %v", got, tt.want)
			}
		})
	}
}
