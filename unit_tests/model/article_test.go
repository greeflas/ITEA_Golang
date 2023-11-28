package model

import (
	"fmt"
	"github.com/google/uuid"
	"testing"
)

func TestNewArticle(t *testing.T) {
	expectedId := uuid.MustParse("4c91ef70-5893-4b1e-96cb-2f7a2ae63593")
	expectedTitle := "This is test article"

	article := NewArticle(expectedId, expectedTitle)

	if article.Id.String() != expectedId.String() {
		t.Errorf("invalid id: got: %s, want: %s", article.Id, expectedId)
	}

	if article.Title != expectedTitle {
		t.Errorf("invalid title: got: %s, want: %s", article.Title, expectedTitle)
	}

	if !article.CreatedAt.Equal(article.UpdatedAt) {
		t.Error("UpdatedAt must be equal to CreatedAt")
	}
}

func TestSetBody(t *testing.T) {
	testCases := map[string]struct {
		bodyValue   string
		errorReason string
	}{
		"empty body": {
			bodyValue:   "",
			errorReason: "cannot set empty body",
		},
		"too short body": {
			bodyValue:   "this is test.",
			errorReason: fmt.Sprintf("cannot set body lower then %d", minBodyLength),
		},
		"too big body": {
			bodyValue:   "this is test, this is test, this is test, this is test, this is test, this is test, this is test, this is test",
			errorReason: fmt.Sprintf("cannot set body grater then %d", maxBodyLength),
		},
	}

	for tn, tc := range testCases {
		t.Run(tn, func(t *testing.T) {
			expectedId := uuid.MustParse("4c91ef70-5893-4b1e-96cb-2f7a2ae63593")
			expectedTitle := "This is test article"
			article := NewArticle(expectedId, expectedTitle)

			err := article.SetBody(tc.bodyValue)

			if err == nil {
				t.Error(tc.errorReason)
			}
		})
	}
}

func TestPublishSuccess(t *testing.T) {
	expectedId := uuid.MustParse("4c91ef70-5893-4b1e-96cb-2f7a2ae63593")
	expectedTitle := "This is test article"

	article := NewArticle(expectedId, expectedTitle)
	article.Body = "This is test body."
	err := article.Publish()

	if err != nil {
		t.Errorf("unexpected error: %s", err)
	}

	if !article.UpdatedAt.After(article.CreatedAt) {
		t.Error("UpdatedAt must be grater then CreatedAt")
	}
}

func TestPublishWithoutBody(t *testing.T) {
	expectedId := uuid.MustParse("4c91ef70-5893-4b1e-96cb-2f7a2ae63593")
	expectedTitle := "This is test article"

	article := NewArticle(expectedId, expectedTitle)
	err := article.Publish()

	if err == nil {
		t.Error("cannot publish article without body")
	}

	if !article.UpdatedAt.Equal(article.CreatedAt) {
		t.Error("UpdatedAt must be equal to CreatedAt")
	}
}
