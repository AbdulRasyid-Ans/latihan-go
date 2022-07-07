package main

import (
	"fmt"
	"strconv"
	"strings"
)

/*
	2 slice of map
		- Slice of Product
			- id
			- name
			- stock
		- Slice of transaksi
			- id
			- username
			- id_product
			- amount

	jika id_product valid maka tampilkan detail nya
*/

// ===== My Answer =====
func main() {
	var products = []map[string]string{
		{
			"id":    "1",
			"name":  "Susu",
			"stock": "10",
		},
		{
			"id":    "2",
			"name":  "Kopi",
			"stock": "5",
		},
	}

	var transactions = []map[string]string{
		{
			"id":         "1",
			"username":   "anon1",
			"id_product": "1",
			"amount":     "3",
		},
		{
			"id":         "2",
			"username":   "anon2",
			"id_product": "3",
			"amount":     "4",
		},
	}

	for _, transaction := range transactions {
		for _, product := range products {
			if transaction["id_product"] == product["id"] {
				productStock, _ := strconv.Atoi(product["stock"])
				tAmmount, _ := strconv.Atoi(transaction["amount"])
				delete(transaction, "id_product")
				fmt.Printf("product_name : %s \n", product["name"])
				fmt.Printf("product_stock : %d \n", (productStock - tAmmount))
				break
			}
		}
		for keyTransaction, valueTransaction := range transaction {
			_ = keyTransaction
			_ = valueTransaction
			fmt.Printf("%s: %s \n", keyTransaction, valueTransaction)
		}
		fmt.Println(strings.Repeat("=", 20))
	}

}

// ===== Instructor Answer =====
/*
func main() {
	var products = []map[string]string{
		{
			"id":    "1",
			"name":  "Susu",
			"stock": "10",
		},
		{
			"id":    "2",
			"name":  "Kopi",
			"stock": "5",
		},
	}

	var transactions = []map[string]string{
		{
			"id":         "1",
			"username":   "anon1",
			"id_product": "1",
			"amount":     "3",
		},
		{
			"id":         "2",
			"username":   "anon2",
			"id_product": "3",
			"amount":     "4",
		},
	}

	for _, t := range transactions {
		pName := ""
		pStock := ""
		isExist := false

		for _, p := range products {
			if t["id_product"] == p["id"] {
				pName = p["name"]
				pStock = p["stock"]
				isExist = true
				break
			}
		}

		fmt.Println("username :", t["username"])
		fmt.Println("id :", t["id"])
		fmt.Println("amount :", t["amount"])
		if isExist {
			fmt.Println("product_name :", pName)
			fmt.Println("product_stock :", pStock)
		} else {
			fmt.Println("id_product :", t["id_product"])
		}

		fmt.Println(strings.Repeat("=", 20))
	}

}
*/
