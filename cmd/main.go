package main

import (
	"GameShopPlatformServer/configs"
	"GameShopPlatformServer/database"
	"GameShopPlatformServer/internal/controllers"
	"GameShopPlatformServer/internal/repository"
	"GameShopPlatformServer/internal/routes"
	"GameShopPlatformServer/internal/service"
	"context"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func init() {
	if err := godotenv.Load("./configs/.env"); err != nil {
		log.Print("No .env file found")
	}
}

func main() {
	dbConf := configs.LoadDBConfig()
	serverConf := configs.LoadServerConfig()

	dbConn, err := database.DBInit(dbConf)
	if err != nil {
		log.Fatalf("error connecting to db: %s", err.Error())
	}
	defer dbConn.Close(context.Background())

	database.RunMigrations(dbConn)

	shopRepo := repository.NewShopRepository(dbConn)
	shopServ := service.NewShopService(shopRepo)
	shopCont := controllers.NewShopController(shopServ)
	shopRout := routes.NewShopRouteController(shopCont)

	gin.SetMode(gin.ReleaseMode)
	router := gin.Default()

	shopRout.ShopRoutes(router)

	if err := router.Run(serverConf.Host + ":" + serverConf.Port); err != nil {
		log.Fatalf("error starting server: %s", err.Error())
	}
}
