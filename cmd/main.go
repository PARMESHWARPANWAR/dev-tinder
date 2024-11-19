package main

import (
	"time"

	route "github.com/PARMESHWARPANWAR/dev-tinder/api/routes"
	"github.com/PARMESHWARPANWAR/dev-tinder/bootstrap"
	"github.com/gin-gonic/gin"
)

func main() {

	app := bootstrap.App()

	env := app.Env

	db := app.Mongo.Database(env.DBName)
	defer app.CloseDBConnection()

	timeout := time.Duration(env.ContextTimeout) * time.Second

	gin := gin.Default()

	route.Setup(env, timeout, db, gin)

	gin.Run(env.ServerAddress)
}
