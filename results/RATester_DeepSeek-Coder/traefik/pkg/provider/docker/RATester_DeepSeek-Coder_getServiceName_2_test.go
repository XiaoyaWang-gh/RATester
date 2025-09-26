package docker

import (
	"fmt"
	"testing"
)

func TestGetServiceName_2(t *testing.T) {
	type args struct {
		container dockerData
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "Test with valid data",
			args: args{
				container: dockerData{
					ID:          "1234567890",
					ServiceName: "test_service",
					Labels: map[string]string{
						labelDockerComposeService: "test_service",
					},
				},
			},
			want: "test_service",
		},
		{
			name: "Test with valid data and multiple labels",
			args: args{
				container: dockerData{
					ID:          "1234567890",
					ServiceName: "test_service",
					Labels: map[string]string{
						labelDockerComposeService:    "test_service",
						"com.docker.compose.project": "test_project",
					},
				},
			},
			want: "test_service_test_project",
		},
		{
			name: "Test with invalid data",
			args: args{
				container: dockerData{
					ID:          "1234567890",
					ServiceName: "test_service",
					Labels: map[string]string{
						"invalid_label": "invalid_value",
					},
				},
			},
			want: "test_service",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			if got := getServiceName(tt.args.container); got != tt.want {
				t.Errorf("getServiceName() = %v, want %v", got, tt.want)
			}
		})
	}
}
