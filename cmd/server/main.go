package main

import "github.com/jabeedhexanovamedia/echo-starter/internal/server"

func main() {
	e := server.New()
	e.Logger.Fatal(e.Start(":8080"))
}
