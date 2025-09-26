package client

import (
	"fmt"
	"testing"

	"github.com/valyala/fasthttp"
)

func TestAddDatas_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	formData := &FormData{
		Args: fasthttp.AcquireArgs(),
	}

	data := map[string][]string{
		"key1": {"value1", "value2"},
		"key2": {"value3", "value4"},
	}

	formData.AddDatas(data)

	if len(formData.Args.Peek("key1")) == 0 || len(formData.Args.Peek("key2")) == 0 {
		t.Error("AddDatas failed")
	}

	fasthttp.ReleaseArgs(formData.Args)
}
