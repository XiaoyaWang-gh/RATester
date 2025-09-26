package http

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/traefik/traefik/v3/pkg/rules"
)

func TestAddRule_2(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	// CONTEXT_0
	m := &matchersTree{}

	// CONTEXT_1
	funcs := make(matcherFuncs)
	funcs["method"] = func(m *matchersTree, value ...string) error {
		return nil
	}

	// CONTEXT_2
	rule := &rules.Tree{
		Matcher: "method",
		Value:   []string{"GET"},
	}

	// CONTEXT_3
	err := m.addRule(rule, funcs)
	require.NoError(t, err)

	// CONTEXT_4
	rule = &rules.Tree{
		Matcher: "and",
		RuleLeft: &rules.Tree{
			Matcher: "method",
			Value:   []string{"GET"},
		},
		RuleRight: &rules.Tree{
			Matcher: "path",
			Value:   []string{"/"},
		},
	}

	// CONTEXT_5
	err = m.addRule(rule, funcs)
	require.NoError(t, err)

	// CONTEXT_6
	funcs["path"] = func(m *matchersTree, value ...string) error {
		return nil
	}

	// CONTEXT_7
	err = m.addRule(rule, funcs)
	require.NoError(t, err)
}
