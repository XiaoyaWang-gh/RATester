package hexec

import (
	"fmt"
	"os/exec"
	"testing"
)

func TestStdinPipe_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	c := &cmdWrapper{
		name: "echo",
		c:    exec.Command("echo", "hello"),
	}
	w, err := c.StdinPipe()
	if err != nil {
		t.Fatal(err)
	}
	defer w.Close()
	if _, err := w.Write([]byte("world")); err != nil {
		t.Fatal(err)
	}
	if err := c.Run(); err != nil {
		t.Fatal(err)
	}
	if c.outerr.String() != "hello world\n" {
		t.Fatalf("unexpected output: %q", c.outerr.String())
	}
}
