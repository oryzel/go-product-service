package product

import (
	"context"

	"github.com/oryzel/go-product-service/models"
)

// Usecase represent the product's usecases
type Usecase interface {
	Fetch(ctx context.Context, cursor string, num int64) ([]*models.Product, string, error)
	GetByID(ctx context.Context, id int64) (*models.Product, error)
	Update(ctx context.Context, ar *models.Product) error
	GetByName(ctx context.Context, title string) (*models.Product, error)
	Insert(context.Context, *models.Product) error
	Delete(ctx context.Context, id int64) error
}
