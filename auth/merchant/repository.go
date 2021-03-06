package merchant

import (
	"github.com/models"
	"golang.org/x/net/context"
)

type Repository interface {
	GetMerchantByName(ctx context.Context,merchantName string)(*models.Merchant,error)
	Fetch(ctx context.Context, cursor string, num int64) (res []*models.Merchant, nextCursor string, err error)
	GetByID(ctx context.Context, id string) (*models.Merchant, error)
	GetByMerchantEmail(ctx context.Context, merchantEmail string) (*models.Merchant, error)
	Update(ctx context.Context, ar *models.Merchant) error
	Insert(ctx context.Context, a *models.Merchant) error
	Delete(ctx context.Context, id string, deleted_by string) error
	Count(ctx context.Context) (int, error)
	List(ctx context.Context, limit, offset int,search string) ([]*models.Merchant, error)
	GetMerchantTransport(ctx context.Context) ([]*models.MerchantTransport, error)
	GetMerchantExperience(ctx context.Context) ([]*models.MerchantExperience, error)
}
