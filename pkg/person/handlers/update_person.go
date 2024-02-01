package person

import (
	"net/http"
	"strconv"

	"github.com/Doni-githu/todo-app"
	"github.com/gin-gonic/gin"
)


type UpdatePersonRequestBody struct {
	Name        string `json:"name"`
	Surname     string `json:"surname"`
	Patronymic  string `json:"patronymic,omitempty"`
	Age         int    `json:"age,omitempty"`
	Gender      string `json:"gender,omitempty"`
	Nationality string `json:"nationality,omitempty"`
}


func (h *handler) UpdatePerson(ctx *gin.Context) {
	id, errId := strconv.Atoi(ctx.Param("id"))
	if errId != nil {
		ctx.AbortWithError(http.StatusBadRequest, errId)
		return
	}
	var body todo.UpdateInput
	err := body.Validate()
	if err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}
	if err := ctx.BindJSON(&body); err != nil  {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	person, err2 := h.s.PersonService.UpdatePerson(body, id)
	if err2 != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	ctx.JSON(http.StatusOK, person)
}