package repository

import "github.com/jackc/pgx/v5"

type ShopRepository struct {
	db *pgx.Conn
}

func NewShopRepository(db *pgx.Conn) *ShopRepository {
	return &ShopRepository{db: db}
}
