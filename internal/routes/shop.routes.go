package routes

import (
	"GameShopPlatformServer/internal/controllers"

	"github.com/gin-gonic/gin"
)

type ShopRouteController struct {
	shopCont *controllers.ShopController
}

func NewShopRouteController(cont *controllers.ShopController) *ShopRouteController {
	return &ShopRouteController{shopCont: cont}
}

func (rc *ShopRouteController) ShopRoutes(router *gin.Engine) {
	auth := router.Group("auth")
	{
		auth.POST("/sign-in", rc.shopCont.SignIn)
		auth.POST("/sign-up", rc.shopCont.SignUp)
	}
}
