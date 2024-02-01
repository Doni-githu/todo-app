package main

import (
	"github.com/Doni-githu/todo-app/pkg/common/db"
	"github.com/Doni-githu/todo-app/pkg/common/middlewares"
	person "github.com/Doni-githu/todo-app/pkg/person/handlers"
	"github.com/Doni-githu/todo-app/pkg/person/repository"
	"github.com/Doni-githu/todo-app/pkg/person/services"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

func main() {
	viper.SetConfigFile("./pkg/common/envs/.env")
	viper.ReadInConfig()

	port := viper.Get("PORT").(string)
	dbUrl := viper.Get("DB_URL").(string)

	gin.SetMode("debug")
	r := gin.Default()
	db := db.Init(dbUrl)
	repo := repository.NewRepository(db)
	s := services.NewService(repo)
	person.RegisterRoutes(r, s)
	r.Use(middlewares.CORSMiddleware())
	
	r.Run(":" +port)
}
