package router

import (
    "github.com/julienschmidt/httprouter"
    "exercise/webservice/controller"
    "exercise/webservice/exception"
    "exercise/webservice/middleware"
)

func NewRouter() *httprouter.Router {
    authController := controller.NewAuthController()
    playerController := controller.NewPlayerController()

    router := httprouter.New()

    router.POST("/api/register", authController.Register)
    router.POST("/api/login", authController.Login)

    router.GET("/api/players", middleware.AuthMiddlewareGroup(playerController.GetAll))
    router.GET("/api/players/:id", middleware.AuthMiddlewareGroup(playerController.FindById))
    router.POST("/api/players", middleware.AuthMiddlewareGroup(playerController.Create))
    router.PUT("/api/players/:id", middleware.AuthMiddlewareGroup(playerController.Update))
    router.DELETE("/api/players/:id", middleware.AuthMiddlewareGroup(playerController.Delete))

    router.PanicHandler = exception.ErrorHandler

    return router
}