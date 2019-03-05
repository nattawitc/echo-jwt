package main

import (
	"fmt"
	"io/ioutil"

	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func main() {

	e := echo.New()
	b, err := ioutil.ReadFile("jwtRSA.key.pub")
	if err != nil {
		fmt.Println("cannot read rsa public key file jwtRSA.key.pub:", err)
		return
	}
	key, err := jwt.ParseRSAPublicKeyFromPEM(b)
	if err != nil {
		fmt.Println("cannot parse public key:", err)
		return
	}

	cfg := middleware.JWTConfig{
		SigningKey:    key,
		SigningMethod: "RS256",
	}
	e.Use(middleware.JWTWithConfig(cfg))

	// Routes
	e.GET("/", hello)

	// Start server
	e.Logger.Fatal(e.Start(":9000"))
}

func hello(c echo.Context) error {
	token := c.Get("user")
	return c.JSON(200, token)
}
