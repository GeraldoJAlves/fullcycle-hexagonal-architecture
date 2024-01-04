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
	defer stmt.Close()

	var product application.Product

	err = stmt.QueryRowContext(ctx, id).Scan(&product.ID, &product.Name, &product.Price, &product.Status)
	if err != nil {
		return nil, err
	}

	return &product, nil
}

func (p ProductDb) Save(product application.ProductInterface) (application.ProductInterface, error) {
	ctx := context.Background()

	var hasProduct bool
	p.db.QueryRowContext(ctx, "SELECT true FROM products WHERE id = ?", product.GetID()).Scan(&hasProduct)

	if hasProduct {
		return p.update(product)
	} else {
		return p.create(product)
	}
}

func (p ProductDb) create(product application.ProductInterface) (application.ProductInterface, error) {
	ctx := context.Background()

	stmt, err := p.db.PrepareContext(ctx, "INSERT INTO products VALUES (?, ?, ?, ?)")
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	_, err = stmt.ExecContext(ctx, product.GetID(), product.GetName(), product.GetPrice(), product.GetStatus())
	if err != nil {
		return nil, err
	}

	return product, nil
}

func (p ProductDb) update(product application.ProductInterface) (application.ProductInterface, error) {
	ctx := context.Background()

	stmt, err := p.db.PrepareContext(ctx, "UPDATE products SET name=?, price=?, status=? WHERE id=?")

	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	_, err = stmt.ExecContext(ctx, product.GetName(), product.GetPrice(), product.GetStatus(), product.GetID())
	if err != nil {
		return nil, err
	}

	return product, nil
}
