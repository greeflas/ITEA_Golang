package main

import (
	"context"
	"errors"
	"flag"
	"fmt"
	"github.com/cucumber/godog"
	"github.com/greeflas/itea_golang/cmd"
)

type CommandStepHandler struct {
	registry *cmd.Registry
}

func NewCommandStepHandler(registry *cmd.Registry) *CommandStepHandler {
	return &CommandStepHandler{registry: registry}
}

func (h *CommandStepHandler) RegisterSteps(ctx *godog.ScenarioContext) {
	ctx.Step(`^I run "([^"]*)" command$`, h.iRunCommand)
	ctx.Step(`^I run "([^"]*)" command with flags:$`, h.iRunCommandWithFlags)
}

func (h *CommandStepHandler) iRunCommand(cmdName string) error {
	command := h.registry.FindCommand(cmdName)

	if command == nil {
		return errors.New("command not found")
	}

	return command.Run(context.Background())
}

func (h *CommandStepHandler) iRunCommandWithFlags(cmdName string, table *godog.Table) error {
	command := h.registry.FindCommand(cmdName)

	if command == nil {
		return errors.New("command not found")
	}

	var args []string
	cellsCount := len(table.Rows[0].Cells)

	for cellIndex := 0; cellIndex < cellsCount; cellIndex++ {
		for rowIndex, row := range table.Rows {
			value := row.Cells[cellIndex].Value

			if rowIndex == 0 {
				value = "--" + value
			}

			args = append(args, value)
		}
	}

	c, ok := command.(FlagsAwareCommand)
	if !ok {
		return fmt.Errorf("command %q hasn't flags", cmdName)
	}

	flagSet := flag.NewFlagSet(cmdName, flag.ExitOnError)
	c.InitFlags(flagSet)

	if err := flagSet.Parse(args); err != nil {
		return err
	}

	return command.Run(context.Background())
}
