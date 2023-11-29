package provider

import (
	"context"
	"fmt"
	"github.com/greeflas/itea_lessons/model"
	"github.com/greeflas/itea_lessons/repository"
)

type DatabaseQuestionsProvider struct {
	questionRepository *repository.QuestionRepository
	answerRepository   *repository.AnswerRepository
}

func NewDatabaseQuestionsProvider(questionRepository *repository.QuestionRepository, answerRepository *repository.AnswerRepository) *DatabaseQuestionsProvider {
	return &DatabaseQuestionsProvider{questionRepository: questionRepository, answerRepository: answerRepository}
}

func (p *DatabaseQuestionsProvider) GetQuestions() ([]*model.Question, error) {
	ctx := context.Background()

	questions, err := p.questionRepository.GetAll(ctx)
	if err != nil {
		return nil, fmt.Errorf("GetQuestions: cannot get questions: %w", err)
	}

	for _, question := range questions {
		answers, err := p.answerRepository.GetAllByQuestionId(ctx, question.Id)
		if err != nil {
			return nil, fmt.Errorf("GetQuestions: cannot get answers: %w", err)
		}

		question.Answers = answers
	}

	return questions, nil
}
