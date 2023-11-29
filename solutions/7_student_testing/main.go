package main

import (
	"context"
	goenv "github.com/greeflas/itea_goenv"
	"github.com/greeflas/itea_lessons/provider"
	"github.com/greeflas/itea_lessons/repository"
	"github.com/jackc/pgx/v5"
	"os"
)

func main() {
	if err := goenv.LoadEnv("./.env"); err != nil {
		panic(err)
	}

	ctx := context.Background()

	conn, err := pgx.Connect(ctx, os.Getenv("DB_CONN_STR"))
	if err != nil {
		panic(err)
	}
	defer conn.Close(ctx)

	questionRepository := repository.NewQuestionRepository(conn)
	answerRepository := repository.NewAnswerRepository(conn)
	questionsProvider := provider.NewDatabaseQuestionsProvider(questionRepository, answerRepository)
	//questionsProvider := provider.NewHardCodedQuestionsProvider()

	studentTestRunner := NewStudentTestRunner(questionsProvider)

	if err := studentTestRunner.Run(); err != nil {
		panic(err)
	}
}
