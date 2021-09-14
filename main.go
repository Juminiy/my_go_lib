package main

import (
	_ "github.com/Juminiy/my_go_lib/web/config"
	"github.com/Juminiy/my_go_lib/web/router"
	"github.com/Juminiy/my_go_lib/web/server"
	"log"
	"os"
)

func main() {
	app := server.InitServer()
	router.ApiBase(app)
	log.Fatalln(app.Listen(os.Getenv("serverPort")))
}
