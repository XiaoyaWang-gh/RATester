package tcp

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/traefik/traefik/v3/pkg/rules"
)

func TestAddRule_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	// Given
	m := &matchersTree{}
	rule := &rules.Tree{
		Matcher: "and",
		RuleLeft: &rules.Tree{
			Matcher: "contains",
			Value:   []string{"foo"},
		},
		RuleRight: &rules.Tree{
			Matcher: "contains",
			Value:   []string{"bar"},
		},
	}
	funcs := make(matcherFuncs)
	funcs["contains"] = func(m *matchersTree, value ...string) error {
		return nil
	}

	// When
	err := m.addRule(rule, funcs)

	// Then
	require.NoError(t, err)
	require.NotNil(t, m.left)
	require.NotNil(t, m.right)
	require.Equal(t, "and", m.operator)
}
