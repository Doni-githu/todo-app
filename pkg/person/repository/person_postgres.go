package repository

import (
	"github.com/Doni-githu/todo-app"
	"github.com/Doni-githu/todo-app/pkg/common/models"
	"gorm.io/gorm"
)

type PersonPostgres struct {
	db *gorm.DB
}

func NewPersonPostgres(db *gorm.DB) *PersonPostgres {
	return &PersonPostgres{db: db}
}

func (r *PersonPostgres) GetAll(page, limit int) ([]models.Person, error) {
	var people []models.Person
	if result := r.db.Find(&people).Limit(limit).Offset(page); result.Error != nil {
		return nil, result.Error
	}
	return nil, nil
}

func (r *PersonPostgres) GetPerson(userId int) (models.Person, error) {
	var person models.Person

	if result := r.db.First(&person); result.Error != nil {
		return models.Person{}, result.Error
	}
	return person, nil
}

func (r *PersonPostgres) AddPerson(body models.AddPerson) (models.Person, error) {
	var person models.Person
	if result := r.db.Create(&person); result.Error != nil {
		return models.Person{}, result.Error
	}
	return models.Person{}, nil
}

func (r *PersonPostgres) UpdatePerson(userId int, input todo.UpdateInput) (models.Person, error) {
	oldPersonData, err := r.GetPerson(userId)
	if err != nil {
		return models.Person{}, err
	}
	name := input.Name
	surname := input.Surname
	patronymic := input.Patronymic
	gender := input.Gender
	nationality := input.Nationality
	age := input.Age
	if name == nil {
		name = &oldPersonData.Name
	}
	if surname == nil {
		surname = &oldPersonData.Surname
	}
	if patronymic == nil {
		patronymic = &oldPersonData.Patronymic
	}
	if gender == nil {
		gender = &oldPersonData.Gender
	}
	if nationality == nil {
		nationality = &oldPersonData.Nationality
	}
	if age == nil {
		age = &oldPersonData.Age
	}
	person := models.Person{
		Name:        *name,
		Surname:     *surname,
		Patronymic:  *patronymic,
		Age:         *age,
		Gender:      *gender,
		Nationality: *nationality,
	}
	result := r.db.Table("people").Where("id = ?", userId).Updates(person)
	if result.Error != nil {
		return models.Person{}, result.Error
	}
	return person, nil
}

func (r *PersonPostgres) DeletePerson(userId int) (string, error) {
	p, err := r.GetPerson(userId)
	if err != nil {
		return "", err
	}

	if result := r.db.Delete(p); result.Error != nil {
		return "", result.Error
	}
	return "Delete Person Successfyly", nil
}

func (r *PersonPostgres) GetPersonWithNameAndSurnameAndPatronymic(name, surname, patronymic string) (models.Person, error) {
	person := models.Person{}
	result := r.db.Where(&models.Person{Name: name, Surname: surname, Patronymic: patronymic}).First(&person)
	return person, result.Error
}
