package todo

import "errors"

type AddPersonRequestBody struct {
	Name       *string `json:"name"`
	Surname    *string `json:"surname"`
	Patronymic *string `json:"patronymic"`
}

func (i AddPersonRequestBody) Validate() error {

	if i.Name == nil || i.Surname == nil || i.Patronymic == nil {
		return errors.New("add structure has no values")
	}

	return nil
}

type NationalityResponse struct {
	Count   int    `json:"count"`
	Name    string `json:"name"`
	Country []struct {
		CountryID   string  `json:"country_id"`
		Probability float64 `json:"probability"`
	} `json:"country"`
}

type AgeResponse struct {
	Count int    `json:"count"`
	Name  string `json:"name"`
	Age   int    `json:"age"`
}

type GenderResponse struct {
	Count       int    `json:"count"`
	Name        string `json:"name"`
	Gender      string `json:"gender"`
	Probability int    `json:"probability"`
}

type UpdateInput struct {
	Name        *string `json:"name"`
	Surname     *string `json:"surname"`
	Patronymic  *string `json:"patronymic,omitempty"`
	Age         *int    `json:"age,omitempty"`
	Gender      *string `json:"gender,omitempty"`
	Nationality *string `json:"nationality,omitempty"`	
}