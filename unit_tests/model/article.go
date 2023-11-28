package model

import (
	"errors"
	"fmt"
	"github.com/google/uuid"
	"time"
)

const (
	minBodyLength = 30
	maxBodyLength = 50
)

type Article struct {
	Id        uuid.UUID
	Title     string
	Body      string
	CreatedAt time.Time
	UpdatedAt time.Time
}

func NewArticle(id uuid.UUID, title string) *Article {
	now := time.Now()

	return &Article{
		Id:        id,
		Title:     title,
		CreatedAt: now,
		UpdatedAt: now,
	}
}

func (a *Article) SetBody(body string) error {
	if body == "" {
		return errors.New("body must not be empty")
	}

	if len(body) < minBodyLength {
		return fmt.Errorf("body must be grater then %d", minBodyLength)
	}

	if len(body) > maxBodyLength {
		return fmt.Errorf("body must be lower then %d", maxBodyLength)
	}

	a.Body = body
	a.UpdatedAt = time.Now()

	return nil
}

func (a *Article) Publish() error {
	if a.Body == "" {
		return errors.New("body is required to publish article")
	}

	a.UpdatedAt = time.Now()

	return nil
}
