package model

import "github.com/google/uuid"

type Answer struct {
	Id        uuid.UUID `json:"id"`
	Text      string    `json:"text"`
	IsCorrect bool      `json:"is_correct"`
}

func NewAnswer(text string, isCorrect bool) *Answer {
	return &Answer{
		Text:      text,
		IsCorrect: isCorrect,
	}
}
