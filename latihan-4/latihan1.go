package main

/*
	2 buah struct
		- struct product
			- id
			- name
			- stock
		- struct transaction
			- id		int
			- amount	int
			- product	product

	Buat method pake Interface. List Method :
	Product :
		- CreateProduct => bool
		- GetProducts => slice of product

	Transaction :
		- CreateTrx => bool
			- Product
			- amount int
		- GetTrxByIdTrx => Trx
			- id_trx, int
*/

import (
	"fmt"
	"strconv"
)

type productMethod interface {
	create() bool
	getList() []Product
}

type Product struct {
	Id    int
	Name  string
	Stock int
}

var products []Product

// create implements productMethod
func (p *Product) create() bool {
	products = append(products, Product{
		p.Id,
		p.Name,
		p.Stock,
	})
	return true
}

// getList implements productMethod
func (*Product) getList() []Product {
	return products
}

type Transaction struct {
	Id     int
	Amount int
	Product
}

func newProduct(name string, stock int) productMethod {
	return &Product{
		Id:    len(products) + 1,
		Stock: stock,
		Name:  name,
	}
}

func listProduct() productMethod {
	return &Product{}
}

func main() {
	bulkCreateProduct()
	lp := listProduct().getList()
	fmt.Println(lp)
}

func bulkCreateProduct() {
	dataProduct := []map[string]string{
		{
			"name":  "Test",
			"stock": "2",
		},
		{
			"name":  "Tost",
			"stock": "10",
		},
		{
			"name":  "Tust",
			"stock": "11",
		},
	}
	for _, produk := range dataProduct {
		s, _ := strconv.Atoi(produk["stock"])
		p := newProduct(produk["name"], s)
		p.create()
	}
}
