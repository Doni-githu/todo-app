package main

import (
	"github.com/Doni-githu/todo-app/pkg/common/db"
	person "github.com/Doni-githu/todo-app/pkg/person/handlers"
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
	r.ForwardedByClientIP = true
	r.SetTrustedProxies([]string{"127.0.0.1"})
	h := db.Init(dbUrl)
	person.RegisterRoutes(r, h)

	r.Run(port)
}
