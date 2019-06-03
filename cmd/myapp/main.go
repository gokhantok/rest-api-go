package main

import(
	"fmt"
	"github.com/labstack/echo"
)

func main(){
	fmt.Println("Welcome to the server")

	InitialMigration() 

	e := echo.New()

	e.GET("/", yallo)
	e.GET("events", getEvents)
	e.POST("events", addEvent)
	e.Start(":8000")
	

}