package controllers

import (
	"GameShopPlatformServer/internal/service"

	"github.com/gin-gonic/gin"
)

type ShopController struct {
	shopServ *service.ShopService
}

func NewShopController(serv *service.ShopService) *ShopController {
	return &ShopController{shopServ: serv}
}

func (sc *ShopController) SignIn(c *gin.Context) {

}

func (sc *ShopController) SignUp(c *gin.Context) {

}
