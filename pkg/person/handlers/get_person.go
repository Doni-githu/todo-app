package person

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func (h *handler) GetPerson(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}
	person, err2 := h.s.PersonService.GetPerson(id)
	if err2 != nil {
		ctx.AbortWithError(http.StatusInternalServerError, err2)
		return
	}
	ctx.JSON(http.StatusOK, person)
}
