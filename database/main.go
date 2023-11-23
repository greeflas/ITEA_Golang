package main

import (
	"context"
	"fmt"
	"github.com/google/uuid"
	"github.com/greeflas/itea_lessons/model"
	"github.com/jackc/pgx/v5"
	"time"
)

func main() {
	ctx := context.Background()

	// postgres://username:user_pass@server_ip:server_port/db_name
	connStr := "postgres://postgres:pass@localhost:5432/lessons"
	conn, err := pgx.Connect(ctx, connStr)
	if err != nil {
		panic(err)
	}
	defer conn.Close(ctx)

	user := model.NewUser(
		uuid.New(),
		randomEmail(),
	)

	_, err = conn.Exec(
		ctx,
		"INSERT INTO users (id, email, is_active) VALUES ($1, $2, $3)",
		user.Id,
		user.Email,
		user.IsActive,
	)
	if err != nil {
		panic(err)
	}

	row := conn.QueryRow(
		ctx,
		"SELECT id, email, is_active FROM users WHERE id = $1",
		user.Id,
	)

	user = new(model.User)

	if err := row.Scan(&user.Id, &user.Email, &user.IsActive); err != nil {
		panic(err)
	}

	user.PrintInfo()
}

func randomEmail() string {
	return fmt.Sprintf("tester+%d@example.com", time.Now().Unix())
}
