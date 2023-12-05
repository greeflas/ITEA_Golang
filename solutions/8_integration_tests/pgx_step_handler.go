package main

import (
	"context"
	"fmt"
	"github.com/cucumber/godog"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgxpool"
	"strings"
)

type PgxStepHandler struct {
	conn *pgxpool.Pool
}

func NewPgxStepHandler(conn *pgxpool.Pool) *PgxStepHandler {
	return &PgxStepHandler{conn: conn}
}

func (h *PgxStepHandler) RegisterSteps(ctx *godog.ScenarioContext) {
	ctx.Step(`^I see in "([^"]*)" table:$`, h.iSeeInTable)
}

func (h *PgxStepHandler) iSeeInTable(tableName string, table *godog.Table) error {
	var columnNames []string

	for _, headerCell := range table.Rows[0].Cells {
		columnNames = append(columnNames, headerCell.Value)
	}

	for _, row := range table.Rows[1:] {
		var expectedValues []string

		for _, cell := range row.Cells {
			expectedValues = append(expectedValues, cell.Value)
		}

		r := h.conn.QueryRow(
			context.Background(),
			fmt.Sprintf("SELECT (%s) FROM %s", strings.Join(columnNames, ","), tableName),
		)

		var unknownValues any
		if err := r.Scan(&unknownValues); err != nil {
			return err
		}

		for i, unknownValue := range unknownValues.([]any) {
			switch value := unknownValue.(type) {
			case string:
				if value != expectedValues[i] {
					return fmt.Errorf("got: %s, want: %s", value, expectedValues[i])
				}
			case uuid.UUID:
				if value.String() != expectedValues[i] {
					return fmt.Errorf("got: %s, want: %s", value, expectedValues[i])
				}
			}
		}
	}

	return nil
}
