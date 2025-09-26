package commands

import (
	"context"
	"fmt"
	"testing"

	"github.com/bep/simplecobra"
)

func TestnewModCommands_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	rootCmd := &rootCommand{}
	modCmd := newModCommands()
	modCmd.r = rootCmd

	// Test Commands() function
	commands := modCmd.Commands()
	if len(commands) == 0 {
		t.Error("Commands() function should return a non-empty list of commands")
	}

	// Test Init() function
	err := modCmd.Init(&simplecobra.Commandeer{})
	if err != nil {
		t.Errorf("Init() function should not return an error, but got: %v", err)
	}

	// Test Name() function
	name := modCmd.Name()
	if name != "mod" {
		t.Errorf("Name() function should return 'mod', but got: %s", name)
	}

	// Test PreRun() function
	err = modCmd.PreRun(&simplecobra.Commandeer{}, &simplecobra.Commandeer{})
	if err != nil {
		t.Errorf("PreRun() function should not return an error, but got: %v", err)
	}

	// Test Run() function
	err = modCmd.Run(context.Background(), &simplecobra.Commandeer{}, []string{})
	if err != nil {
		t.Errorf("Run() function should not return an error, but got: %v", err)
	}
}
