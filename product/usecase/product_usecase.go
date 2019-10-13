package usecase

import (
	"context"
	"time"

	"github.com/oryzel/go-product-service/models"
	"github.com/oryzel/go-product-service/product"
)

type productUsecase struct {
	productRepo    product.Repository
	contextTimeout time.Duration
}

// NewProductUsecase will create new an productUsecase object representation of product.Usecase interface
func NewProductUsecase(p product.Repository, timeout time.Duration) product.Usecase {
	return &productUsecase{
		productRepo:    p,
		contextTimeout: timeout,
	}
}

func (p *productUsecase) Fetch(c context.Context, cursor string, num int64) ([]*models.Product, string, error) {
	if num == 0 {
		num = 10
	}

	ctx, cancel := context.WithTimeout(c, p.contextTimeout)
	defer cancel()

	listProduct, nextCursor, err := p.productRepo.Fetch(ctx, cursor, num)
	if err != nil {
		return nil, "", err
	}

	return listProduct, nextCursor, nil
}

func (p *productUsecase) GetByID(c context.Context, id int64) (*models.Product, error) {

	ctx, cancel := context.WithTimeout(c, p.contextTimeout)
	defer cancel()

	res, err := p.productRepo.GetByID(ctx, id)
	if err != nil {
		return nil, err
	}

	return res, nil
}

func (p *productUsecase) GetByName(c context.Context, title string) (*models.Product, error) {

	ctx, cancel := context.WithTimeout(c, p.contextTimeout)
	defer cancel()
	res, err := p.productRepo.GetByName(ctx, title)
	if err != nil {
		return nil, err
	}

	return res, nil
}

func (p *productUsecase) Insert(c context.Context, m *models.Product) error {

	ctx, cancel := context.WithTimeout(c, p.contextTimeout)
	defer cancel()
	existedProduct, _ := p.GetByName(ctx, m.Name)
	if existedProduct != nil {
		return models.ErrConflict
	}

	err := p.productRepo.Insert(ctx, m)
	if err != nil {
		return err
	}
	return nil
}

func (p *productUsecase) Update(c context.Context, pr *models.Product) error {

	ctx, cancel := context.WithTimeout(c, p.contextTimeout)
	defer cancel()

	pr.UpdatedAt = time.Now()
	return p.productRepo.Update(ctx, pr)
}

func (p *productUsecase) Delete(c context.Context, id int64) error {
	ctx, cancel := context.WithTimeout(c, p.contextTimeout)
	defer cancel()
	existedArticle, err := p.productRepo.GetByID(ctx, id)
	if err != nil {
		return err
	}
	if existedArticle == nil {
		return models.ErrNotFound
	}
	return p.productRepo.Delete(ctx, id)
}
