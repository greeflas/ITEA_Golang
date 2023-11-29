package repository

import (
	"context"
	"fmt"
	"github.com/google/uuid"
	"github.com/greeflas/itea_lessons/model"
	"github.com/jackc/pgx/v5"
)

type AnswerRepository struct {
	conn *pgx.Conn
}

func NewAnswerRepository(conn *pgx.Conn) *AnswerRepository {
	return &AnswerRepository{conn: conn}
}

func (r *AnswerRepository) GetAllByQuestionId(ctx context.Context, questionId uuid.UUID) ([]*model.Answer, error) {
	sql := "SELECT id, text, is_correct FROM answers WHERE question_id = $1"
	rows, err := r.conn.Query(ctx, sql, questionId)
	if err != nil {
		return nil, fmt.Errorf("pgx error: %w", err)
	}

	var answers []*model.Answer

	for rows.Next() {
		answer := new(model.Answer)
		if err := rows.Scan(&answer.Id, &answer.Text, &answer.IsCorrect); err != nil {
			return nil, fmt.Errorf("pgx error: %w", err)
		}

		answers = append(answers, answer)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("pgx error: %w", err)
	}

	return answers, nil
}
