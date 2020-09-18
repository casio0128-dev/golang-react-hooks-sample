package main

import (
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func main() {
	port := flag.String("port", "1323", "this args is port number.")

	e := echo.New()

	e.Use(middleware.CORS())
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	root, err := os.Getwd()

	if err == nil {
		fmt.Println(root)
	} else {
		log.Fatal(err)
		return
	}

	e.Static("/", "../client/build")

	e.Logger.Fatal(e.Start(fmt.Sprintf(":%s", *port)))
	fmt.Println(root)
}
