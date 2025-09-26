package fiber

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"testing"
	"time"
)

func TestGet_4(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	r := New()
	r.Get("/", func(c Ctx) error {
		return c.SendString("Hello, World!")
	})
	go func() {
		err := r.Listen(":3000")
		if err != nil {
			t.Error(err)
		}
	}()
	time.Sleep(time.Second)
	res, err := http.Get("http://localhost:3000")
	if err != nil {
		t.Error(err)
	}
	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		t.Error(err)
	}
	if string(body) != "Hello, World!" {
		t.Errorf("Expected %s, got %s", "Hello, World!", string(body))
	}
}
