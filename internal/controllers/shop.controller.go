package controllers

import (
	"GameShopPlatformServer/internal/service"
	"GameShopPlatformServer/models"
	cerr "GameShopPlatformServer/pkg/custom-errors"
	"GameShopPlatformServer/pkg/jwts"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type ShopController struct {
	shopServ *service.ShopService
}

func NewShopController(serv *service.ShopService) *ShopController {
	return &ShopController{shopServ: serv}
}

func (sc *ShopController) SignIn(c *gin.Context) {
	token := c.GetHeader("Authorization")

	_, err := jwts.ParseToken(token)
	if err == nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": cerr.ErrAlreadyAuthorized.Error()})
		return
	}

	var user models.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	userId, err := sc.shopServ.GetShopRepository().GetAuthUserData(&user)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": cerr.ErrUserNotFound.Error()})
		return
	}

	roles, err := sc.shopServ.GetShopRepository().GetUsersRoles(userId)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": cerr.ErrUserNotFound.Error()})
		return
	}

	userIDstr := strconv.Itoa(userId)
	tokenNew, err := jwts.CreateToken(userIDstr, roles)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": cerr.ErrClaims.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"token": tokenNew})
}

func (sc *ShopController) SignUp(c *gin.Context) {
	token := c.GetHeader("Authorization")

	_, err := jwts.ParseToken(token)
	if err == nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": cerr.ErrAlreadyAuthorized.Error()})
		return
	}

	var user models.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err = sc.shopServ.GetShopRepository().CreateUser(&user)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": cerr.ErrAlreadyRegistered.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "user created"})
}
