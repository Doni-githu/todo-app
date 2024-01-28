package person

import (
	"fmt"
	"net/http"

	"github.com/Doni-githu/todo-app/pkg/common/models"
	"github.com/gin-gonic/gin"
)

func (h *handler) GetPerson(ctx *gin.Context) {
	id := ctx.Param("id")
	var person models.Person

	fmt.Printf("id: %v\n", id)

	if result := h.db.First(&person, id); result.Error != nil {
		ctx.AbortWithError(http.StatusNotFound, result.Error)
	}

	ctx.JSON(http.StatusOK, person)
}
