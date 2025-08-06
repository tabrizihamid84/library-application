package main

import (
	"net/http"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/tabrizihamid84/library-application/api/route"
	"github.com/tabrizihamid84/library-application/bootstrap"
)

func main() {
	app := bootstrap.App()

	env := app.Env

	db := app.Mongo.Database(env.DBName)
	defer app.CloseDB()

	timeout := time.Duration(env.ContextTimeout) * time.Second

	echo := echo.New()

	echo.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"https://labstack.com", "http://localhost:3000"},
		AllowMethods: []string{http.MethodGet, http.MethodPut, http.MethodPost, http.MethodDelete},
	}))

	route.Setup(env, timeout, db, echo)

	echo.GET("/", HealthCheck)
	// echo.GET("/swagger/*", echoSwagger.WrapHandler)

	echo.Logger.Fatal(echo.Start(env.ServerAddress))
}

func HealthCheck(c echo.Context) error {
	return c.JSON(http.StatusOK, map[string]interface{}{
		"data": "Server is up and running",
	})

}
