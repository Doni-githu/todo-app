package services

import (
	"github.com/Doni-githu/todo-app"
	"github.com/Doni-githu/todo-app/pkg/common/models"
	"github.com/Doni-githu/todo-app/pkg/person/repository"
)

type PersonService struct {
	repo     repository.Person
	external *ExternalService
}

func NewPersonService(repo repository.Person, external *ExternalService) *PersonService {
	return &PersonService{repo: repo, external: external}
}

func (s *PersonService) GetPeople(page, limit int) ([]models.Person, error) {
	return s.repo.GetAll(page, limit)
}

func (s *PersonService) GetPerson(userId int) (models.Person, error) {
	return s.repo.GetPerson(userId)
}

func (s *PersonService) AddPerson(body models.AddPerson) (models.Person, error) {
	nat, _ := s.external.GetNationality(body.Name)
	age, _ := s.external.GetAge(body.Name)
	gender, _ := s.external.GetGender(body.Name)
	person := models.AddPerson{
		Name:        body.Name,
		Surname:     body.Surname,
		Patronymic:  body.Patronymic,
		Age:         age,
		Gender:      gender,
		Nationality: nat,
	}
	return s.repo.AddPerson(person)
}

func (s *PersonService) UpdatePerson(body todo.UpdateInput, userId int) (models.Person, error) {
	return s.repo.UpdatePerson(userId, body)
}

func (s *PersonService) DeletePerson(userId int) (string, error) {
	return s.repo.DeletePerson(userId)
}

func (s *PersonService) GetPersonWithNameAndSurnameAndPatronymic(name, surname, patronymic string) (models.Person, error) {
	return s.repo.GetPersonWithNameAndSurnameAndPatronymic(name, surname, patronymic)
}
