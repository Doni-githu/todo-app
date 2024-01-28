package person

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type handler struct {
	db *gorm.DB
}

func RegisterRoutes(r *gin.Engine, db *gorm.DB) {
	h := &handler{
		db: db,
	}

	routes := r.Group("/people")
	routes.POST("/", h.AddPerson)
	routes.GET("/", h.GetPeople)
	routes.GET("/{id}/", h.GetPerson)
	routes.PUT("/{id}/", h.UpdatePerson)
	routes.DELETE("/{id}/", h.DeletePerson)
}