package main

import (
	"database/sql"
	db2 "github.com/codeedu/go-hexagonal/adapters/db"
	"github.com/codeedu/go-hexagonal/application"
	_ "github.com/mattn/go-sqlite3"
)

func main() {
	db, _ := sql.Open("sqlite3", ":memory:")
	productDbAdapter := db2.NewProductDb(db)

	productService := application.NewProductService(productDbAdapter)
	product, _ := productService.Create("Product Test", 10.0)

	productService.Enable(product)
}
