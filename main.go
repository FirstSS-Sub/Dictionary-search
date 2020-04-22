package main

import (
	"fmt"
	"github.com/FirstSS-Sub/Dictionary-search/handler"
	"github.com/FirstSS-Sub/Dictionary-search/injector"
	"github.com/labstack/echo"
)

func main() {
	fmt.Println("sever start")
	dictHandler := injector.InjectDictHandler()
	e := echo.New()
	handler.InitRouting(e, dictHandler)
	e.Logger.Fatal(e.Start(":8080"))
}
