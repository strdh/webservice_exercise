package main

import (
    "fmt"
    "net/http"
    // "github.com/go-playground/validator/v10"
    _ "github.com/go-sql-driver/mysql"
    // "exercise/webservice/app/config"
    // "exercise/webservice/repository"
    // "exercise/webservice/service"
    // "exercise/webservice/controller"
    "exercise/webservice/helper"
    "exercise/webservice/router"
)

func main() {
    // db := app.NewDB()
    // validate := validator.New()
    // playerRepository := repository.NewPlayerRepository()
    // playerService := service.NewPlayerService(playerRepository, db, validate)
    // playerController := controller.NewPlayerController()

    router := router.NewRouter()

    server := http.Server{
        Addr: "localhost:8080",
        Handler: router,
    }

    picture := `       
                ,_     _
                |\\_,-~/
                / _  _ |    ,--.
                (  @  @ )   / ,-'
                \\  _T_/-._( (
                /         ` + "`" + `. \\
                |         _  \\ 
                \\ \\ ,  /    |
                || |-_\\__   /
                ((_/` + "`" + `(____,-'
                ----------------
-----------------SERVER STARTED-----------------
                ----------------
`

	// print the picture
	fmt.Println(picture)

    err := server.ListenAndServe()
    helper.PanicIfError(err)
}