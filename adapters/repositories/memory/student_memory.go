package memory

import (
	"github.com/Kriwn/go-sd-2/entities"
	"github.com/Kriwn/go-sd-2/repositories"
)

type studentMemory struct {
	students []entities.Student
}


func NewStudentMemory() repositories.StudentRepository {
	return &studentMemory{
		students: []entities.Student{
			{
				ID: "651xxxxx",
				Name: "Kritt Wongwandee",
			},
		},
	}
}

func (s *studentMemory) GetALL() []entities.Student {
	return s.students
}
