package proxy

import (
	"fmt"
	"testing"

	"github.com/valyala/fasthttp"
)

func TestWithClient_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	// arrange
	var cli *fasthttp.Client
	// act
	WithClient(cli)
	// assert
}
