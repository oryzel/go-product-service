package repository

import (
	"context"
	"database/sql"
	"encoding/base64"
	"fmt"
	"time"

	"github.com/sirupsen/logrus"

	"github.com/oryzel/go-product-service/models"
	"github.com/oryzel/go-product-service/product"
)

const (
	timeFormat = "2006-01-02T15:04:05.999Z07:00" // reduce precision from RFC3339Nano as date format
)

type mysqlProductRepository struct {
	Conn *sql.DB
}

// NewMysqlArticleRepository will create an object that represent the article.Repository interface
func NewMysqlArticleRepository(Conn *sql.DB) product.Repository {
	return &mysqlProductRepository{Conn}
}

func (m *mysqlProductRepository) fetch(ctx context.Context, query string, args ...interface{}) ([]*models.Product, error) {
	rows, err := m.Conn.QueryContext(ctx, query, args...)
	if err != nil {
		logrus.Error(err)
		return nil, err
	}

	defer func() {
		err := rows.Close()
		if err != nil {
			logrus.Error(err)
		}
	}()

	result := make([]*models.Product, 0)
	for rows.Next() {
		t := new(models.Product)
		err = rows.Scan(
			&t.ID,
			&t.Name,
			&t.Sku,
			&t.Category,
			&t.Type,
			&t.CreatedAt,
			&t.UpdatedAt,
		)

		if err != nil {
			logrus.Error(err)
			return nil, err
		}

		result = append(result, t)
	}

	return result, nil
}

func (m *mysqlProductRepository) Fetch(ctx context.Context, cursor string, num int64) ([]*models.Product, string, error) {
	query := `SELECT id, name, sku, category, type, created_at, updated_at
  						FROM products ORDER BY created_at LIMIT ? `

	res, err := m.fetch(ctx, query, num)
	if err != nil {
		return nil, "", err
	}

	nextCursor := ""
	if len(res) == int(num) {
		nextCursor = EncodeCursor(res[len(res)-1].CreatedAt)
	}

	return res, nextCursor, err
}

func (m *mysqlProductRepository) GetByID(ctx context.Context, id int64) (res *models.Product, err error) {
	query := `SELECT id, name, sku, category, type, created_at, updated_at
  						FROM products WHERE ID = ?`

	list, err := m.fetch(ctx, query, id)
	if err != nil {
		return nil, err
	}

	if len(list) > 0 {
		res = list[0]
	} else {
		return nil, models.ErrNotFound
	}

	return
}

func (m *mysqlProductRepository) GetByName(ctx context.Context, title string) (res *models.Product, err error) {
	query := `SELECT id, name , sku, type, category, updated_at, created_at
  						FROM products WHERE name = ?`

	list, err := m.fetch(ctx, query, title)
	if err != nil {
		return
	}

	if len(list) > 0 {
		res = list[0]
	} else {
		return nil, models.ErrNotFound
	}
	return
}

func (m *mysqlProductRepository) Insert(ctx context.Context, p *models.Product) error {
	query := `INSERT  products SET name=? , sku=? , category=?, type=? , created_at=?, updated_at=?`
	stmt, err := m.Conn.PrepareContext(ctx, query)
	if err != nil {
		return err
	}

	res, err := stmt.ExecContext(ctx, p.Name, p.Sku, p.Category, p.Type, p.CreatedAt, p.UpdatedAt)
	if err != nil {
		return err
	}

	lastID, err := res.LastInsertId()
	if err != nil {
		return err
	}

	p.ID = lastID
	return nil
}

func (m *mysqlProductRepository) Update(ctx context.Context, p *models.Product) error {
	query := `UPDATE products set name=?, sku=?, category=?, type=?, updated_at=? WHERE ID = ?`

	stmt, err := m.Conn.PrepareContext(ctx, query)
	if err != nil {
		return nil
	}

	res, err := stmt.ExecContext(ctx, p.Name, p.Sku, p.Category, p.Type, p.UpdatedAt, p.ID)
	if err != nil {
		return err
	}
	affect, err := res.RowsAffected()
	if err != nil {
		return err
	}
	if affect != 1 {
		err = fmt.Errorf("Weird  Behaviour. Total Affected: %d", affect)

		return err
	}

	return nil
}

func (m *mysqlProductRepository) Delete(ctx context.Context, id int64) error {
	query := "DELETE FROM products WHERE id = ?"

	stmt, err := m.Conn.PrepareContext(ctx, query)
	if err != nil {
		return err
	}

	res, err := stmt.ExecContext(ctx, id)
	if err != nil {

		return err
	}

	rowsAfected, err := res.RowsAffected()
	if err != nil {
		return err
	}

	if rowsAfected != 1 {
		err = fmt.Errorf("Weird  Behaviour. Total Affected: %d", rowsAfected)
		return err
	}

	return nil
}

// DecodeCursor will decode cursor from user for mysql
func DecodeCursor(encodedTime string) (time.Time, error) {
	byt, err := base64.StdEncoding.DecodeString(encodedTime)
	if err != nil {
		return time.Time{}, err
	}

	timeString := string(byt)
	t, err := time.Parse(timeFormat, timeString)

	return t, err
}

// EncodeCursor will encode cursor from mysql to user
func EncodeCursor(t time.Time) string {
	timeString := t.Format(timeFormat)

	return base64.StdEncoding.EncodeToString([]byte(timeString))
}
