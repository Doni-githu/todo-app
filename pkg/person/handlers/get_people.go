package person

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func (h *handler) GetPeople(ctx *gin.Context) {
	page, err3:= strconv.Atoi(ctx.Query("page"))
	limit, err2:= strconv.Atoi(ctx.Query("limit"))
	if err3 != nil {
		page = 0
	}
	if err2 != nil {
		limit = 10
	}

	people, err := h.s.PersonService.GetPeople(page, limit)
	if err != nil {
		ctx.AbortWithError(http.StatusNotFound, err)
		return
	}
	ctx.JSON(http.StatusOK, people)
}