package main

import (
	"time"

	"github.com/MaogeJing/gohome/api/route"
	"github.com/MaogeJing/gohome/bootstrap"
	"github.com/gin-gonic/gin"
)

func main() {
	app := bootstrap.App()
	env := app.Env

	mysql_db := app.MySQL
	defer app.CloseDBConnection()

	timeout := time.Duration(env.ContextTimeout) * time.Second
	gin := gin.Default()

	route.Setup(env, timeout, mysql_db, gin)

	gin.Run(env.ServerAddress)

}
