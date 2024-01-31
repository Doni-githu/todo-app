package services

import (
	"github.com/Doni-githu/todo-app"
	"github.com/Doni-githu/todo-app/pkg/common/models"
	"github.com/Doni-githu/todo-app/pkg/person/repository"
)

type ExternalServices interface {
	GetAge(name string) (int, error)
	GetNationality(name string) (string, error)
	GetGender(name string) (string, error)
}

type PersonServices interface {
	GetPeople(page, limit int) ([]models.Person, error)
	GetPerson(userId int) (models.Person, error)
	AddPerson(user models.AddPerson) (models.Person, error)
	UpdatePerson(body todo.UpdateInput, userId int) (models.Person, error)
	DeletePerson(userId int) (string, error)
	GetPersonWithNameAndSurnameAndPatronymic(name, surname, patronymic string) *models.Person
}

type Service struct {
	ExternalServices
	PersonService
}

func NewService(repo *repository.Repository) *Service {
	external := NewExternalService()

	return &Service{
		ExternalServices: external,
		PersonService:    *NewPersonService(repo.Person, external),
	}
}
