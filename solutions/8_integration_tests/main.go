package main

import (
	"context"
	"flag"
	"github.com/greeflas/itea_golang/cmd"
	"github.com/greeflas/itea_golang/repository"
	pgxuuid "github.com/jackc/pgx-gofrs-uuid"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	"os"
)

type FlagsAwareCommand interface {
	InitFlags(flags *flag.FlagSet)
}

func main() {
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

	cmdName := os.Args[1]
	command := commandRegistry.FindCommand(cmdName)
	if command == nil {
		panic("command not found")
	}

	flagSet := flag.NewFlagSet(os.Args[0], flag.ExitOnError)

	if c, ok := command.(FlagsAwareCommand); ok {
		c.InitFlags(flagSet)
	}

	if err := flagSet.Parse(os.Args[2:]); err != nil {
		panic(err)
	}

	if err := command.Run(ctx); err != nil {
		panic(err)
	}
}
