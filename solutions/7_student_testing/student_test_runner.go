package main

import (
	"fmt"
	"github.com/greeflas/itea_lessons/model"
)

type QuestionsProvider interface {
	GetQuestions() ([]*model.Question, error)
}

type StudentTestRunner struct {
	questionsProvider QuestionsProvider
}

func NewStudentTestRunner(questionsProvider QuestionsProvider) *StudentTestRunner {
	return &StudentTestRunner{questionsProvider: questionsProvider}
}

func (r *StudentTestRunner) Run() error {
	questions, err := r.questionsProvider.GetQuestions()
	if err != nil {
		return fmt.Errorf("run: cannot get questions from provider: %w", err)
	}

	correctAnswersCount, incorrectAnswersCount, err := r.askQuestions(questions)
	if err != nil {
		return err
	}

	r.printResult(correctAnswersCount, incorrectAnswersCount)

	return nil
}

func (r *StudentTestRunner) askQuestions(questions []*model.Question) (correctAnswersCount int, incorrectAnswersCount int, err error) {
	for _, question := range questions {
		answerOptions := r.createAnswerOptions(question)

		r.printQuestionWithAnswerOptions(question, answerOptions)

		studentsAnswer, err := r.askAnswer(answerOptions)
		if err != nil {
			return 0, 0, fmt.Errorf("askQuestions: cannot get students answer: %w", err)
		}

		if question.IsCorrectAnswer(studentsAnswer) {
			correctAnswersCount++
		} else {
			incorrectAnswersCount++
		}
	}

	return
}

func (r *StudentTestRunner) printResult(correctAnswersCount, incorrectAnswersCount int) {
	fmt.Printf("\nWell done!\n")
	fmt.Printf("Number of correct answers is: %d\n", correctAnswersCount)
	fmt.Printf("Number of incorrect answers is: %d\n", incorrectAnswersCount)
}

func (r *StudentTestRunner) createAnswerOptions(question *model.Question) map[string]*model.Answer {
	var answerOptions = make(map[string]*model.Answer)

	for i, answer := range question.Answers {
		variantIndex := string(rune('a' + i))
		answerOptions[variantIndex] = answer
	}

	return answerOptions
}

func (r *StudentTestRunner) printQuestionWithAnswerOptions(question *model.Question, options map[string]*model.Answer) {
	fmt.Printf("\n%s\n", question.Text)

	for option, answer := range options {
		fmt.Printf("%s) %s\n", option, answer.Text)
	}
}

func (r *StudentTestRunner) askAnswer(answerOptions map[string]*model.Answer) (*model.Answer, error) {
	for {
		fmt.Print("\nYour variant > ")

		var studentsAnswer string
		if _, err := fmt.Scan(&studentsAnswer); err != nil {
			return nil, fmt.Errorf("askAnswer: scan error: %w", err)
		}

		answer, ok := answerOptions[studentsAnswer]
		if ok {
			return answer, nil
		}

		fmt.Println("There is no such variant of answer for this question. Try again.")
	}
}
