package main

import (
	"context"
	"github.com/cucumber/godog"
	"github.com/greeflas/itea_golang/cmd"
	"github.com/greeflas/itea_golang/repository"
	pgxuuid "github.com/jackc/pgx-gofrs-uuid"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	"testing"
)

func TestFeatures(t *testing.T) {
	ctx := context.Background()

	connStr := "postgres://postgres:pass@localhost:5432/lessons"
	dbconfig, err := pgxpool.ParseConfig(connStr)
	if err != nil {
		panic(err)
	}
	dbconfig.AfterConnect = func(ctx context.Context, conn *pgx.Conn) error {
		pgxuuid.Register(conn.TypeMap())
		return nil
	}

	dbpool, err := pgxpool.NewWithConfig(ctx, dbconfig)
	if err != nil {
		panic(err)
	}
	defer dbpool.Close()

	articleRepository := repository.NewArticleRepository(dbpool)
	createArticleCommand := cmd.NewCreateArticleCommand(articleRepository)
	updateArticleCommand := cmd.NewUpdateArticleCommand(articleRepository)
	commandRegistry := cmd.NewRegistry(createArticleCommand, updateArticleCommand)

	suite := godog.TestSuite{
		Name: "Articles agency",
		Options: &godog.Options{
			Format: "pretty",
			Paths: []string{
				"features",
			},
		},
		ScenarioInitializer: func(ctx *godog.ScenarioContext) {
			ctx.Before(func(ctx context.Context, s *godog.Scenario) (context.Context, error) {
				// TODO: cleanup database here
				return ctx, nil
			})

			commandStepHandler := NewCommandStepHandler(commandRegistry)
			commandStepHandler.RegisterSteps(ctx)

			pgxStepHandler := NewPgxStepHandler(dbpool)
			pgxStepHandler.RegisterSteps(ctx)
		},
	}

	if suite.Run() != 0 {
		t.Fatal("non-zero status returned, failed to run feature test")
	}
}
