package gorm

import (
	"demo/gorm/models"
	"testing"
)

func TestSchool_Add(t *testing.T) {
	var teacher models.Teacher
	teacher.Name ="Tom"
	var students []*models.Student
	students = append(students, &models.Student{Name: "Jesse"})
	students = append(students, &models.Student{Name: "Jack"})
	teacher.Student = students
	errs := (*School)(nil).Add(&teacher)
	if len(errs) > 0{
		t.Fatal(errs)
	}
}