package db_test

import (
	"database/sql"
	"log"
	"testing"

	"github.com/geraldojalves/fullcycle-hexagonal-architecture/internal/adapters/db"
	"github.com/geraldojalves/fullcycle-hexagonal-architecture/internal/application"
	_ "github.com/mattn/go-sqlite3"
	"github.com/stretchr/testify/require"
)

var sqliteDb *sql.DB

func setUp() {
	sqliteDb, _ = sql.Open("sqlite3", ":memory:")
	createTable(sqliteDb)
}

func createTable(db *sql.DB) {
	table := `CREATE TABLE products (
		id string,
		name string,
		price float,
		status string
	)`

	stmt, err := db.Prepare(table)

	if err != nil {
		log.Fatal(err.Error())
	}

	stmt.Exec()
}

func insertProduct(db *sql.DB, id string, name string, price float64, status string) {
	stmt, err := db.Prepare(`INSERT INTO products values(?, ?, ?, ?)`)
	if err != nil {
		log.Fatal(err)
	}

	stmt.Exec(id, name, price, status)
}

func TestProductDb_Get(t *testing.T) {
	setUp()
	defer sqliteDb.Close()

	productDb := db.NewProductDb(sqliteDb)

	insertProduct(sqliteDb, "uuid-1", "ball", 2.1, "enabled")

	p, err := productDb.Get("uuid-1")

	require.Nil(t, err)
	require.Equal(t, "uuid-1", p.GetID())
	require.Equal(t, "ball", p.GetName())
	require.Equal(t, 2.1, p.GetPrice())
	require.Equal(t, "enabled", p.GetStatus())
}

func TestProductDb_Save(t *testing.T) {
	setUp()
	defer sqliteDb.Close()

	productDb := db.NewProductDb(sqliteDb)
	product := application.NewProduct("ball", 1.99)

	_, err := productDb.Save(product)

	require.Nil(t, err)

	productResult, err := productDb.Get(product.GetID())
	require.Nil(t, err)
	require.Equal(t, product.GetID(), productResult.GetID())
	require.Equal(t, product.GetName(), productResult.GetName())
	require.Equal(t, product.GetPrice(), productResult.GetPrice())
	require.Equal(t, product.GetStatus(), productResult.GetStatus())

	product.Name = "ball 2"
	product.Price = 2.99
	product.Enable()

	_, err = productDb.Save(product)

	require.Nil(t, err)

	productResult, err = productDb.Get(product.GetID())
	require.Nil(t, err)
	require.Equal(t, product.GetID(), productResult.GetID())
	require.Equal(t, product.GetName(), productResult.GetName())
	require.Equal(t, product.GetPrice(), productResult.GetPrice())
	require.Equal(t, product.GetStatus(), productResult.GetStatus())
}
