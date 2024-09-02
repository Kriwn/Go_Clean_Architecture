package repositories

import "github.com/Kriwn/go-sd-2/entities"


type StudentRepository interface{
	GetALL() []entities.Student
}
