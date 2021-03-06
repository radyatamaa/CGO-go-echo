package promo_merchant

import (
	"context"
	"github.com/models"
)

type Repository interface {
	CountByPromoId(ctx context.Context,promoId string)(int,error)
	GetByMerchantId(ctx context.Context,merchantId string,promoId string)([]*models.PromoMerchant,error)
	Insert(ctx context.Context,pm models.PromoMerchant)error
	DeleteByMerchantId(ctx context.Context,merchantId string,promoId string)error
}
