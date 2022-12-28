package router

import (
    "github.com/julienschmidt/httprouter"
    "exercise/webservice/controller"
    "exercise/webservice/exception"
)

func NewRouter(playerController controller.PlayerController) *httprouter.Router {
    router := httprouter.New()

    router.GET("/api/players", playerController.GetAll)
    router.GET("/api/players/:id", playerController.FindById)
    router.POST("/api/players", playerController.Create)
    router.PUT("/api/players/:id", playerController.Update)
    router.DELETE("/api/players/:id", playerController.Delete)

    router.PanicHandler = exception.ErrorHandler

    return router
}