package db

import (
	"context"
	"database/sql"

	"github.com/geraldojalves/fullcycle-hexagonal-architecture/internal/application"
)

type ProductDb struct {
	db *sql.DB
}

func NewProductDb(db *sql.DB) *ProductDb {
	return &ProductDb{
		db: db,
	}
}

func (p ProductDb) Get(id string) (application.ProductInterface, error) {
	ctx := context.Background()

	stmt, err := p.db.PrepareContext(ctx, "SELECT id, name, price, status FROM products WHERE id = ?")
	if err != nil {
		return nil, err
	}

	var product application.Product

	err = stmt.QueryRowContext(ctx, id).Scan(&product.ID, &product.Name, &product.Price, &product.Status)
	if err != nil {
		return nil, err
	}

	return &product, nil
}
