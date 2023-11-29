package repository

import (
	"context"
	"fmt"
	"github.com/greeflas/itea_lessons/model"
	"github.com/jackc/pgx/v5"
)

type QuestionRepository struct {
	conn *pgx.Conn
}

func NewQuestionRepository(conn *pgx.Conn) *QuestionRepository {
	return &QuestionRepository{conn: conn}
}

func (r *QuestionRepository) GetAll(ctx context.Context) ([]*model.Question, error) {
	rows, err := r.conn.Query(ctx, "SELECT id, text FROM questions")
	if err != nil {
		return nil, fmt.Errorf("pgx error: %w", err)
	}

	var questions []*model.Question

	for rows.Next() {
		question := new(model.Question)
		if err := rows.Scan(&question.Id, &question.Text); err != nil {
			return nil, fmt.Errorf("pgx error: %w", err)
		}

		questions = append(questions, question)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("pgx error: %w", err)
	}

	return questions, nil
}
