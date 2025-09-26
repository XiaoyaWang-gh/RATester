package client

import (
	"fmt"
	"sort"
	"testing"
)

func TestGetRandomPortArr_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	min := 1024
	max := 65535
	ports := getRandomPortArr(min, max)

	if len(ports) != max-min+1 {
		t.Errorf("Expected length of ports is %d, but got %d", max-min+1, len(ports))
	}

	for _, port := range ports {
		if port < min || port > max {
			t.Errorf("Port number %d is out of range [%d, %d]", port, min, max)
		}
	}

	sort.Ints(ports)
	for i := 0; i < len(ports)-1; i++ {
		if ports[i] == ports[i+1] {
			t.Errorf("Ports array is not random")
		}
	}
}
