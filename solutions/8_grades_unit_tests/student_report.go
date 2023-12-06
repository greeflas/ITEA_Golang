package main

import "fmt"

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
