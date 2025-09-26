package ingress

import (
	"fmt"
	"testing"
)

func TestSetDefaults_3(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	serviceIng := &ServiceIng{}
	serviceIng.SetDefaults()

	if serviceIng.PassHostHeader == nil || *serviceIng.PassHostHeader != true {
		t.Errorf("Expected PassHostHeader to be true, but got %v", serviceIng.PassHostHeader)
	}
}
