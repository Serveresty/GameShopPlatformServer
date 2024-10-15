package controllers

import "GameShopPlatformServer/internal/service"

type ShopController struct {
	shopServ *service.ShopService
}

func NewShopController(serv *service.ShopService) *ShopController {
	return &ShopController{shopServ: serv}
}
