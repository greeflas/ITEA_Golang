package main

import (
	"testing"
)

func TestAverageMarkStudentWithMarks(t *testing.T) {
	student := NewStudent("First", "Last")
	report := NewStudentReport(student)

	report.AddMark(NewRate("test", 90))
	report.AddMark(NewRate("test", 80))
	report.AddMark(NewRate("test", 100))

	actual := report.AverageMark()
	expected := 90

	if actual != expected {
		t.Errorf("invalid mark: got: %d, want: %d", actual, expected)
	}
}

func TestAverageMarkStudentWithoutMarks(t *testing.T) {
	student := NewStudent("First", "Last")
	report := NewStudentReport(student)

	actual := report.AverageMark()
	expected := 0

	if actual != expected {
		t.Errorf("invalid mark: got: %d, want: %d", actual, expected)
	}
}
