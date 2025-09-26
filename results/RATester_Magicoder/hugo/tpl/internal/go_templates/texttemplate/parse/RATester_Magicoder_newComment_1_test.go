package parse

import (
	"fmt"
	"testing"
)

func TestnewComment_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	tree := &Tree{}
	pos := Pos(1)
	text := "test comment"
	comment := tree.newComment(pos, text)

	if comment.Pos != pos {
		t.Errorf("Expected Pos to be %v, but got %v", pos, comment.Pos)
	}

	if comment.Text != text {
		t.Errorf("Expected Text to be %s, but got %s", text, comment.Text)
	}

	if comment.NodeType != NodeComment {
		t.Errorf("Expected NodeType to be NodeComment, but got %v", comment.NodeType)
	}

	if comment.tr != tree {
		t.Errorf("Expected tr to be %p, but got %p", tree, comment.tr)
	}
}
