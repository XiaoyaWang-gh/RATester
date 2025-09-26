package proxy

import (
	"fmt"
	"testing"

	"github.com/zeebo/assert"
)

func TestClose_8(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	// given:
	s := &Sock5ModeServer{
		BaseServer: BaseServer{
			id: 1,
		},
	}
	// when:
	err := s.Close()
	// then:
	assert.NoError(t, err)
}
