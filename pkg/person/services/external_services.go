package services

import (
	"encoding/json"
	"net/http"
	"os"

	"github.com/Doni-githu/todo-app"
)

type ExternalService struct {
	ageApi         string
	genderApi      string
	nationalityApi string
}

func findMaxProbabilityCountry(nationality todo.NationalityResponse) (string, float64) {
	maxProbability := nationality.Country[0].Probability
	maxCountryID := nationality.Country[0].CountryID

	for _, country := range nationality.Country {
		if country.Probability > maxProbability {
			maxProbability = country.Probability
			maxCountryID = country.CountryID
		}
	}

	return maxCountryID, maxProbability
}

func NewExternalService() *ExternalService {
	return &ExternalService{
		ageApi:         os.Getenv("AGE_API"),
		genderApi:      os.Getenv("GENDER_API"),
		nationalityApi: os.Getenv("NATIONALITY_API"),
	}
}

func (s *ExternalService) GetAge(name string) (int, error) {
	url := s.ageApi + "/?name=" + name
	res, err := http.Get(url)
	if err != nil {
		return 0, err
	}
	defer res.Body.Close()
	var j todo.AgeResponse
	err = json.NewDecoder(res.Body).Decode(&j)
	if err != nil {
		return 0, err
	}

	return j.Age, nil
}

func (s *ExternalService) GetNationality(name string) (string, error) {
	url := s.nationalityApi + "/?name=" + name
	res, err := http.Get(url)
	if err != nil {
		return "", err
	}
	defer res.Body.Close()
	var j todo.NationalityResponse
	err2 := json.NewDecoder(res.Body).Decode(&j)
	if err2 != nil {
		return "", err2
	}
	country, _ := findMaxProbabilityCountry(j)
	return country, nil
}

func (s *ExternalService) GetGender(name string) (string, error) {
	url := s.genderApi + "/?name=" + name
	res, err := http.Get(url)
	if err != nil {
		return "", err
	}
	defer res.Body.Close()
	var j todo.GenderResponse
	err = json.NewDecoder(res.Body).Decode(&j)
	if err != nil {
		return "", err
	}

	return j.Gender, nil
}
