package person

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func (h *handler) DeletePerson(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	_, err2 := h.s.PersonService.DeletePerson(id)
	if err2 != nil {
		ctx.AbortWithError(http.StatusBadRequest, err2)
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"message": "Person has been deleted",
	})
}
