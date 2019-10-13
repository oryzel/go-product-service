package product

import (
	"context"

	"github.com/oryzel/go-product-service/models"
)

// Repository represent the product's repository contract
type Repository interface {
	Fetch(ctx context.Context, cursor string, num int64) (res []*models.Product, nextCursor string, err error)
	GetByID(ctx context.Context, id int64) (*models.Product, error)
	GetByName(ctx context.Context, title string) (*models.Product, error)
	Update(ctx context.Context, ar *models.Product) error
	Insert(ctx context.Context, a *models.Product) error
	Delete(ctx context.Context, id int64) error
}
