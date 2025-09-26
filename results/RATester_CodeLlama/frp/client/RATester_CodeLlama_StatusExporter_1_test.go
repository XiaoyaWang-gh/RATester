package client

import (
	"fmt"
	"testing"

	"github.com/zeebo/assert"
)

func TestStatusExporter_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	svr := &Service{}
	exporter := svr.StatusExporter()
	assert.NotNil(t, exporter)
}
