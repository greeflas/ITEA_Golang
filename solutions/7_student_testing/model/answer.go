package model

import "github.com/google/uuid"

type Answer struct {
	Id        uuid.UUID
	Text      string
	IsCorrect bool
}

func NewAnswer(text string, isCorrect bool) *Answer {
	return &Answer{
		Text:      text,
		IsCorrect: isCorrect,
	}
}
