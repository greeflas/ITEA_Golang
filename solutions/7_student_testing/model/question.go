package model

import "github.com/google/uuid"

type Question struct {
	Id      uuid.UUID `json:"id"`
	Text    string    `json:"text"`
	Answers []*Answer `json:"answers"`
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
