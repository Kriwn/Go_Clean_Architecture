package usecases

import (
	"github.com/Kriwn/go-sd-2/entities"
	"github.com/Kriwn/go-sd-2/repositories"
)

type StudentUseCase interface {
	GetALLStudents() []entities.Student
}

// service
type studentService struct {
	studentRepo repositories.StudentRepository
}

func NewStudentService(studentRepo repositories.StudentRepository) StudentUseCase {
	return &studentService{
		studentRepo: studentRepo,
	}
}

func (s *studentService) GetALLStudents() []entities.Student {
	return s.studentRepo.GetALL()
}
