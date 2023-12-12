package provider

import (
	"encoding/json"
	"fmt"
	"github.com/greeflas/itea_lessons/model"
	"io"
	"os"
)

type JsonFileQuestionsProvider struct {
	filePath string
}

func NewJsonFileQuestionsProvider(filePath string) *JsonFileQuestionsProvider {
	return &JsonFileQuestionsProvider{filePath: filePath}
}

func (p *JsonFileQuestionsProvider) GetQuestions() ([]*model.Question, error) {
	file, err := os.Open(p.filePath)
	if err != nil {
		return nil, fmt.Errorf("GetQuestions: failed open file: %w", err)
	}

	bytes, err := io.ReadAll(file)
	if err != nil {
		return nil, fmt.Errorf("GetQuestions: failed read file: %w", err)
	}

	var questions []*model.Question

	if err := json.Unmarshal(bytes, &questions); err != nil {
		return nil, fmt.Errorf("GetQuestions: failed unmarshal json: %w", err)
	}

	return questions, nil
}
