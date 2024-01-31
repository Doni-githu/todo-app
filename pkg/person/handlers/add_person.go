package person

import (
	"net/http"

	todo "github.com/Doni-githu/todo-app"
	"github.com/gin-gonic/gin"
)

func (h *handler) AddPerson(ctx *gin.Context) {
	body := todo.AddPersonRequestBody{}

	if err := ctx.BindJSON(&body); err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}
	sameOne := h.s.PersonService.GetPersonWithNameAndSurnameAndPatronymic(*body.Name, *body.Surname, *body.Patronymic)
	if sameOne != nil {
		ctx.JSON(http.StatusCreated, sameOne)
		return
	}
	ctx.JSON(http.StatusOK, body)
}
