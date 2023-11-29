package main

import (
	"fmt"
)

type Answer struct {
	Text      string
	IsCorrect bool
}

func NewAnswer(text string, isCorrect bool) *Answer {
	return &Answer{
		Text:      text,
		IsCorrect: isCorrect,
	}
}

type Question struct {
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

type QuestionsProvider interface {
	GetQuestions() ([]*Question, error)
}

type HardCodedQuestionsProvider struct{}

func NewHardCodedQuestionsProvider() *HardCodedQuestionsProvider {
	return &HardCodedQuestionsProvider{}
}

func (p HardCodedQuestionsProvider) GetQuestions() ([]*Question, error) {
	firstQuestion := NewQuestion("What is structure in GoLang?")
	firstQuestion.Answers = []*Answer{
		NewAnswer("Interface", false),
		NewAnswer("Typed collection of fields", true),
		NewAnswer("Pointer", false),
	}

	secondQuestion := NewQuestion("You can pass the struct to/from a function as")
	secondQuestion.Answers = []*Answer{
		NewAnswer("Value", true),
		NewAnswer("Link", false),
	}

	thirdQuestion := NewQuestion("Is it possible to create a pointer to a struct?")
	thirdQuestion.Answers = []*Answer{
		NewAnswer("Yes", true),
		NewAnswer("No", false),
	}

	questions := []*Question{firstQuestion, secondQuestion, thirdQuestion}

	return questions, nil
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

func (r *StudentTestRunner) askQuestions(questions []*Question) (correctAnswersCount int, incorrectAnswersCount int, err error) {
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

func (r *StudentTestRunner) createAnswerOptions(question *Question) map[string]*Answer {
	var answerOptions = make(map[string]*Answer)

	for i, answer := range question.Answers {
		variantIndex := string(rune('a' + i))
		answerOptions[variantIndex] = answer
	}

	return answerOptions
}

func (r *StudentTestRunner) printQuestionWithAnswerOptions(question *Question, options map[string]*Answer) {
	fmt.Printf("\n%s\n", question.Text)

	for option, answer := range options {
		fmt.Printf("%s) %s\n", option, answer.Text)
	}
}

func (r *StudentTestRunner) askAnswer(answerOptions map[string]*Answer) (*Answer, error) {
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

func main() {
	questionsProvider := NewHardCodedQuestionsProvider()
	studentTestRunner := NewStudentTestRunner(questionsProvider)

	if err := studentTestRunner.Run(); err != nil {
		panic(err)
	}
}
