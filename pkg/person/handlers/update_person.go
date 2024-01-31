package person

import "github.com/gin-gonic/gin"

type UpdatePersonRequestBody struct {
	Name        string `json:"name"`
	Surname     string `json:"surname"`
	Patronymic  string `json:"patronymic,omitempty"`
	Age         int    `json:"age,omitempty"`
	Gender      string `json:"gender,omitempty"`
	Nationality string `json:"nationality,omitempty"`
}


func (h *handler) UpdatePerson(ctx *gin.Context) {
	
}