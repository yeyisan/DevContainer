package main

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func main() {
	fmt.Println("Hello")
	e := echo.New()

	e.HideBanner = true
	e.HidePort = true
	e.Use(middleware.Gzip())
	e.Use(middleware.Secure())
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowHeaders: []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept, echo.HeaderAuthorization},
	}))

	e.GET("/", Redirect)
	e.GET("/home", Home)
	e.GET("/about", About)

	e.Start(":8080")
}

func Redirect(c echo.Context) error {
	return c.JSON(http.StatusOK, map[string]string{"message": "allowed routes are /home and /about"})
}

func Home(c echo.Context) error {
	return c.JSON(http.StatusOK, map[string]string{"message": "welcome to the home"})
}

func About(c echo.Context) error {
	return c.JSON(http.StatusOK, map[string]string{"message": "there is something about me"})
}
