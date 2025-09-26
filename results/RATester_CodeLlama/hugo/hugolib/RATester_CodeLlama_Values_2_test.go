package hugolib

import (
	"fmt"
	"testing"

	"github.com/alecthomas/assert"
)

func TestValues_2(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	t.Parallel()

	// GIVEN
	given := taxonomiesConfig{
		"foo": "foo",
		"bar": "bar",
		"baz": "baz",
	}

	// WHEN
	actual := given.Values()

	// THEN
	assert.Equal(t, []viewName{
		{singular: "foo", plural: "foo", pluralTreeKey: "foo"},
		{singular: "bar", plural: "bar", pluralTreeKey: "bar"},
		{singular: "baz", plural: "baz", pluralTreeKey: "baz"},
	}, actual.views)

	assert.Equal(t, map[string]viewName{
		"foo": {singular: "foo", plural: "foo", pluralTreeKey: "foo"},
		"bar": {singular: "bar", plural: "bar", pluralTreeKey: "bar"},
		"baz": {singular: "baz", plural: "baz", pluralTreeKey: "baz"},
	}, actual.viewsByTreeKey)
}
