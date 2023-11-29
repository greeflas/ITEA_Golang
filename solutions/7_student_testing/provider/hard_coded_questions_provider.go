package provider

import "github.com/greeflas/itea_lessons/model"

type HardCodedQuestionsProvider struct{}

func NewHardCodedQuestionsProvider() *HardCodedQuestionsProvider {
	return &HardCodedQuestionsProvider{}
}

func (p HardCodedQuestionsProvider) GetQuestions() ([]*model.Question, error) {
	firstQuestion := model.NewQuestion("What is structure in GoLang?")
	firstQuestion.Answers = []*model.Answer{
		model.NewAnswer("Interface", false),
		model.NewAnswer("Typed collection of fields", true),
		model.NewAnswer("Pointer", false),
	}

	secondQuestion := model.NewQuestion("You can pass the struct to/from a function as")
	secondQuestion.Answers = []*model.Answer{
		model.NewAnswer("Value", true),
		model.NewAnswer("Link", false),
	}

	thirdQuestion := model.NewQuestion("Is it possible to create a pointer to a struct?")
	thirdQuestion.Answers = []*model.Answer{
		model.NewAnswer("Yes", true),
		model.NewAnswer("No", false),
	}

	questions := []*model.Question{firstQuestion, secondQuestion, thirdQuestion}

	return questions, nil
}
