package models

import (
	"fmt"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestString_21(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	e := &DateField{}
	e.Set(time.Now())
	assert.Equal(t, e.Value().String(), e.String())
}
