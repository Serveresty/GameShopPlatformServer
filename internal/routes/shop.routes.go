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

}
