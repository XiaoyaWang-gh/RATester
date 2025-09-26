package fiber

import (
	"fmt"
	"io/ioutil"
	"os"
	"testing"
)

func TestprintRoutesMessage_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	app := &App{
		config: Config{
			ColorScheme: Colors{
				Blue:   "blue",
				White:  "white",
				Green:  "green",
				Yellow: "yellow",
				Cyan:   "cyan",
				Reset:  "reset",
			},
		},
		stack: [][]*Route{
			{
				{
					Method: "GET",
					Path:   "/test",
					Name:   "test",
					Handlers: []Handler{
						func(c Ctx) error {
							return nil
						},
					},
				},
			},
		},
	}

	old := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w

	app.printRoutesMessage()

	w.Close()
	out, _ := ioutil.ReadAll(r)
	os.Stdout = old

	expected := "method\t| path\t| name\t| handlers\t\n------\t| ----\t| ----\t| --------\t\nblueGET\t| /test\t| cyantest\t| yellowgithub.com/gofiber/fiber/v3.(*App).printRoutesMessage.func1\n"

	if string(out) != expected {
		t.Errorf("Expected: %s, Got: %s", expected, string(out))
	}
}
