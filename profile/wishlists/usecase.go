package wishlists

import (
	"context"
	"github.com/models"
)

type Usecase interface {
	Insert(ctx context.Context, wl *models.WishlistIn, token string) (string, error)
	List(ctx context.Context, token string,	page int, limit int, offset int,expId string) (*models.WishlistOutWithPagination, error)
}
