package visitor

import (
	"fmt"
	"testing"

	"github.com/zeebo/assert"
)

func TestMsgTransporter_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	v := &visitorHelperImpl{}
	assert.NotNil(t, v.MsgTransporter())
}
