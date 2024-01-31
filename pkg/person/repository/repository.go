package repository

import (
	"github.com/Doni-githu/todo-app"
	"github.com/Doni-githu/todo-app/pkg/common/models"
	"gorm.io/gorm"
)


type Person interface {
	GetAll(page, limit int) ([]models.Person, error)
	GetPerson(userId int) (models.Person, error)
	AddPerson(user models.AddPerson) (models.Person, error)
	UpdatePerson(userId int, input todo.UpdateInput) (models.Person, error)
	DeletePerson(userId int) (string, error)
	GetPersonWithNameAndSurnameAndPatronymic(name, surname, patronymic string) *models.Person
}


type Repository struct {
	Person
}


func NewRepository(db *gorm.DB) *Repository {
	return &Repository{
		Person: NewPersonPostgres(db),
	}
}