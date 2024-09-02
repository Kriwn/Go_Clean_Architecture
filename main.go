package main

import (
	"github.com/Kriwn/go-sd-2/adapters/handler/rest"
	"github.com/Kriwn/go-sd-2/adapters/repositories/memory"
	"github.com/Kriwn/go-sd-2/usecases"
	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()

	studentRepo := memory.NewStudentMemory()
	studentService := usecases.NewStudentService(studentRepo)
	studentHandler := rest.NewStudentRestHandler(studentService)

	app.Get("/students" , studentHandler.GetALLStudents)

	// bind port
	app.Listen(":8000")
}
