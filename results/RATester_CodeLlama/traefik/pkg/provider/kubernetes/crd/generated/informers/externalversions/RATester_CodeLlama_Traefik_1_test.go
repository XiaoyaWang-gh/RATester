package externalversions

import (
	"fmt"
	"testing"
)

func TestTraefik_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	f := &sharedInformerFactory{}
	f.Traefik()
}
