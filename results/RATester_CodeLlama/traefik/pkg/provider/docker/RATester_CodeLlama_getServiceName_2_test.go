package docker

import (
	"fmt"
	"testing"

	"github.com/alecthomas/assert"
)

func TestGetServiceName_2(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	var container dockerData
	container.ServiceName = "test"
	container.Labels = make(map[string]string)
	container.Labels[labelDockerComposeProject] = "test"
	container.Labels[labelDockerComposeService] = "test"
	assert.Equal(t, "test_test", getServiceName(container))
}
