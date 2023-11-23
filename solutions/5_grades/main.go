package main

import "fmt"

type Student struct {
	firstName string
	lastName  string
}

func NewStudent(firstName, lastName string) *Student {
	return &Student{
		firstName: firstName,
		lastName:  lastName,
	}
}

func (s Student) FullName() string {
	return s.firstName + " " + s.lastName
}

type Rate struct {
	subject string
	mark    int
}

func NewRate(subject string, mark int) *Rate {
	return &Rate{
		subject: subject,
		mark:    mark,
	}
}

type StudentReport struct {
	student *Student
	marks   []*Rate
}

func NewStudentReport(student *Student) *StudentReport {
	return &StudentReport{
		student: student,
	}
}

func (s *StudentReport) AddMark(mark *Rate) {
	s.marks = append(s.marks, mark)
}

func (s *StudentReport) AverageMark() int {
	if len(s.marks) == 0 {
		return 0
	}

	var sum int
	for _, mark := range s.marks {
		sum += mark.mark
	}

	return sum / len(s.marks)
}

func (s *StudentReport) Print() {
	fmt.Printf("Student`s name: %s\n", s.student.FullName())
	fmt.Printf("\taverage mark is: %d\n", s.AverageMark())
}

func main() {
	sean := NewStudent("Sean", "Doe")

	seanFirstMark := NewRate("Go Intro", 90)
	seanSecondMark := NewRate("Git", 90)
	seanThirdMark := NewRate("Control Flow", 70)

	seanReport := NewStudentReport(sean)
	seanReport.AddMark(seanFirstMark)
	seanReport.AddMark(seanSecondMark)
	seanReport.AddMark(seanThirdMark)
	seanReport.Print()

	ann := NewStudent("Ann", "Doe")

	annFirstMark := NewRate("Go Intro", 90)
	annSecondMark := NewRate("Git", 100)
	annThirdMark := NewRate("Control Flow", 80)

	annReport := NewStudentReport(ann)
	annReport.AddMark(annFirstMark)
	annReport.AddMark(annSecondMark)
	annReport.AddMark(annThirdMark)
	annReport.Print()

	john := NewStudent("John", "Doe")

	johnReport := NewStudentReport(john)
	johnReport.Print()
}
