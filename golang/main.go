package main

import (
	"flag"
	"fmt"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"log"
	"os"
)

func main() {
	port := flag.String("port", "3000", "this args is port number.")

	if root, err := os.Getwd(); err == nil {
		fmt.Println(root)
	} else {
		log.Fatal(err)
		return
	}

	e := echo.New()

	e.Use(middleware.CORS())
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.Static("/", "client/build/")

	e.Logger.Fatal(e.Start(fmt.Sprintf(":%s", *port)))
}