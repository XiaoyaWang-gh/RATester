package hugolib

import (
	"fmt"
	"html/template"
	"testing"
)

func TestInnerDeindent_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	scp := &ShortcodeWithPage{
		Inner: template.HTML(`
			<p>
				<a href="https://example.com">
					<img src="https://example.com/image.png" alt="example">
				</a>
			</p>
		`),
		indentation: "  ",
	}

	expected := template.HTML(`
		<p>
			<a href="https://example.com">
				<img src="https://example.com/image.png" alt="example">
			</a>
		</p>
	`)

	if scp.InnerDeindent() != expected {
		t.Errorf("expected %q, got %q", expected, scp.InnerDeindent())
	}
}
