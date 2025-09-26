package hugolib

import (
	"context"
	"fmt"
	"testing"

	"github.com/alecthomas/assert"
	"github.com/gohugoio/hugo/markup/tableofcontents"
)

func TestHeadingsFiltered_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	p := &pageHeadingsFiltered{
		headings: tableofcontents.Headings{
			{
				Level: 1,
				Title: "Heading 1",
			},
			{
				Level: 2,
				Title: "Heading 2",
			},
			{
				Level: 3,
				Title: "Heading 3",
			},
		},
	}

	headings := p.HeadingsFiltered(context.Background())

	assert.Equal(t, 3, len(headings))
	assert.Equal(t, 1, headings[0].Level)
	assert.Equal(t, "Heading 1", headings[0].Title)
	assert.Equal(t, 2, headings[1].Level)
	assert.Equal(t, "Heading 2", headings[1].Title)
	assert.Equal(t, 3, headings[2].Level)
	assert.Equal(t, "Heading 3", headings[2].Title)
}
