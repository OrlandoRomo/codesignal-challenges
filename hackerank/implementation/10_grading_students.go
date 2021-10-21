package main

import "fmt"

func main() {
	fmt.Println(gradingStudents([]int32{
		4, 73, 67, 38, 33,
	}))
}

func gradingStudents(grades []int32) []int32 {
	approvedGrades := make([]int32, 0)
	var multiple, gradeCondition int32 = 5, 3
	for _, grade := range grades {
		mod := grade % multiple
		substraction := multiple - mod
		if substraction < gradeCondition && grade >= 38 {
			approvedGrades = append(approvedGrades, int32(grade+substraction))
			continue
		}
		approvedGrades = append(approvedGrades, grade)
	}
	return approvedGrades
}
