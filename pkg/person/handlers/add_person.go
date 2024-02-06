package person

import (
	"net/http"

	"github.com/Doni-githu/todo-app/pkg/common/models"
	"github.com/gin-gonic/gin"
)

func (h *handler) AddPerson(ctx *gin.Context) {
	body := models.AddPerson{}

	if err := ctx.BindJSON(&body); err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}
	sameOne, err := h.s.PersonService.GetPersonWithNameAndSurnameAndPatronymic(body.Name, body.Surname, body.Patronymic)
	if err == nil {
		ctx.JSON(http.StatusCreated, sameOne)
		return
	}
	person, err3 := h.s.PersonService.AddPerson(body)
	if err3 != nil {
		ctx.AbortWithError(http.StatusInternalServerError, err3)
		return 
	}
	ctx.JSON(http.StatusOK, person)
}
