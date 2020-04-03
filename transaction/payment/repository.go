package payment

import (
	"context"
	"github.com/models"
)

type Repository interface {
	Insert(ctx context.Context, payment *models.Transaction) (*models.Transaction, error)
}
