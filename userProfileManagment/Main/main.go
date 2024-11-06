package main

import (
	"userProfileManagment/config"
	"userProfileManagment/controller"
	"userProfileManagment/repository"
	"userProfileManagment/router"
	"userProfileManagment/service"

	"github.com/gin-gonic/gin"
)

func main() {
	db, err := config.InitDB()
	if err != nil {
		panic("failed to connect to database")
	}
	routers := gin.Default()
	repo := repository.NewUserRepository(db)
	services := service.NewUserService(repo)
	controll := controller.NewUserController(services)
	router.InitializeRouter(routers, controll)
	routers.Run()
}
