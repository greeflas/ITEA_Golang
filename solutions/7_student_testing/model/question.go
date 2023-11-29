package model

import "github.com/google/uuid"

type Question struct {
	Id      uuid.UUID
	Text    string
	Answers []*Answer
}

func NewQuestion(text string) *Question {
	return &Question{
		Text: text,
	}
}

func (q *Question) IsCorrectAnswer(studentsAnswer *Answer) bool {
	for _, answer := range q.Answers {
		if answer == studentsAnswer && answer.IsCorrect {
			return true
		}
	}

	return false
}
