package service

import "GameShopPlatformServer/internal/repository"

type ShopService struct {
	shopRepo *repository.ShopRepository
}

func NewShopService(repo *repository.ShopRepository) *ShopService {
	return &ShopService{shopRepo: repo}
}

func (ss *ShopService) GetShopRepository() *repository.ShopRepository {
	return ss.shopRepo
}
