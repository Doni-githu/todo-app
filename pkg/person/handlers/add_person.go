package person

import (
	"encoding/json"
	"net/http"

	"github.com/Doni-githu/todo-app/pkg/common/models"
	"github.com/gin-gonic/gin"
)

type AddPersonRequestBody struct {
	Name       string `json:"name"`
	Surname    string `json:"surname"`
	Patronymic string `json:"patronymic"`
}

type Nationality struct {
	Count   int    `json:"count"`
	Name    string `json:"name"`
	Country []struct {
		CountryID   string  `json:"country_id"`
		Probability float64 `json:"probability"`
	} `json:"country"`
}

type Age struct {
	Count int    `json:"count"`
	Name  string `json:"name"`
	Age   int    `json:"age"`
}

type Gender struct {
	Count       int    `json:"count"`
	Name        string `json:"name"`
	Gender      string `json:"gender"`
	Probability int    `json:"probability"`
}

func findMaxProbabilityCountry(nationality Nationality) (string, float64) {
	maxProbability := 0.0
	maxCountryID := ""

	for _, country := range nationality.Country {
		if country.Probability > maxProbability {
			maxProbability = country.Probability
			maxCountryID = country.CountryID
		}
	}

	return maxCountryID, maxProbability
}

func getNationality(name, surname string) (string, error) {
	url := "https://api.nationalize.io/?" + "name=" + name + "&surname=" + surname
	res, err := http.Get(url)
	if err != nil {
		return "", err
	}
	defer res.Body.Close()
	var j Nationality
	err2 := json.NewDecoder(res.Body).Decode(&j)
	if err2 != nil {
		return "", err2
	}
	country, _ := findMaxProbabilityCountry(j)
	return country, nil
}

func getAge(name string) (int, error) {
	url := "https://api.agify.io/?" + "name=" + name
	res, err := http.Get(url)
	if err != nil {
		return 0, err
	}
	defer res.Body.Close()
	var j Age
	err = json.NewDecoder(res.Body).Decode(&j)
	if err != nil {
		return 0, err
	}

	return j.Age, nil
}

func getGender(name string) (string, error) {
	url := "https://api.genderize.io/?" + "name=" + name
	res, err := http.Get(url)
	if err != nil {
		return "", err
	}
	defer res.Body.Close()
	var j Gender
	err = json.NewDecoder(res.Body).Decode(&j)
	if err != nil {
		return "", err
	}

	return j.Gender, nil
}

func (h *handler) AddPerson(ctx *gin.Context) {
	body := AddPersonRequestBody{}

	if err := ctx.BindJSON(&body); err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	var person models.Person
	nat, _ := getNationality(body.Name, body.Surname)
	age, _ := getAge(body.Name)
	gender, _ := getGender(body.Name)
	person.Nationality = nat
	person.Age = age
	person.Gender = gender
	person.Patronymic = body.Patronymic
	person.Surname = body.Surname
	person.Name = body.Name
	
	if result := h.db.Create(&person); result.Error != nil {
		ctx.AbortWithError(http.StatusNotFound, result.Error)
		return
	}

	ctx.JSON(http.StatusOK, &person)
}
