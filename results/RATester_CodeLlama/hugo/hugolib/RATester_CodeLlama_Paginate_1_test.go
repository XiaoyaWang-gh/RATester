package hugolib

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestPaginate_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	p := &pagePaginator{
		source: &pageState{
			pageOutputs: []*pageOutput{
				{
					render: true,
				},
				{
					render: true,
				},
			},
		},
	}

	pager, err := p.Paginate([]any{1, 2, 3, 4, 5, 6, 7, 8, 9, 10})
	require.NoError(t, err)
	require.NotNil(t, pager)
	require.Equal(t, 2, len(pager.Pagers()))
}
