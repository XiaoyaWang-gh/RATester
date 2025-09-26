package session

import (
	"fmt"
	"testing"

	"github.com/zeebo/assert"
)

func TestCfgEnableSidInURLQuery_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	// arrange
	enable := true
	// act
	opt := CfgEnableSidInURLQuery(enable)
	// assert
	assert.NotNil(t, opt)
}
