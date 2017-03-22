package main

import "github.com/labstack/echo"

func main() {
	e := echo.New()
	setupRoutes(e)
	e.Logger.Fatal(e.Start(":8080"))
}

func setupRoutes(e *echo.Echo) {
	slackHandler := SlackHandler{}
	e.POST("/fortune", slackHandler.Webhook)
}

