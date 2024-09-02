package rest

import (
	"github.com/Kriwn/go-sd-2/usecases"
	"github.com/gofiber/fiber/v2"
)

type StudentRestHandler interface {
	GetALLStudents(c *fiber.Ctx) error
}

type studentResthHandler struct {
	studentUseCase usecases.StudentUseCase
}


func NewStudentRestHandler(studentUseCase usecases.StudentUseCase) StudentRestHandler {
	return &studentResthHandler{
		studentUseCase: studentUseCase,
	}
}

func (s *studentResthHandler) GetALLStudents(c *fiber.Ctx) error{

	students := s.studentUseCase.GetALLStudents()

	return c.Status(fiber.StatusOK).JSON(students)
}
