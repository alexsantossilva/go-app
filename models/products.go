package models

import (
	"fmt"

	"github.com/alexsantossilva/go-app/db"
)

type Product struct {
	Name, Description string
	Price             float64
	Id, Quantity      int
}

func FindAll() []Product {
	db := db.ConectDb()
	findAllProducts, err := db.Query("select * from products order by id ASC")
	if err != nil {
		panic(err.Error())
	}

	product := Product{}
	products := []Product{}

	for findAllProducts.Next() {
		var id, quantity int
		var name, description string
		var price float64

		err = findAllProducts.Scan(&id, &name, &description, &price, &quantity)
		if err != nil {
			panic(err.Error())
		}

		product.Id = id
		product.Name = name
		product.Description = description
		product.Price = price
		product.Quantity = quantity

		products = append(products, product)

	}

	defer db.Close()

	return products
}

func InsertProduct(name, description string, price float64, quantity int) {
	db := db.ConectDb()

	insert, err := db.Prepare("insert into products (name, description, price, quantity) values ($1, $2, $3, $4)")
	if err != nil {
		panic(err.Error())
	}

	insert.Exec(name, description, price, quantity)
	defer db.Close()
}

func Delete(id string) {
	db := db.ConectDb()

	delete, err := db.Prepare("delete from products where id = $1")
	if err != nil {
		panic(err.Error())
	}

	delete.Exec(id)
	defer db.Close()
}

func FindById(id string) Product {
	db := db.ConectDb()

	productQuery, err := db.Query("select * from products where id = $1", id)
	if err != nil {
		panic(err.Error())
	}

	product := Product{}
	for productQuery.Next() {
		var id, quantity int
		var name, description string
		var price float64

		err = productQuery.Scan(&id, &name, &description, &price, &quantity)
		if err != nil {
			panic(err.Error())
		}

		product.Id = id
		product.Name = name
		product.Description = description
		product.Price = price
		product.Quantity = quantity
	}
	defer db.Close()
	return product
}

func Update(id int, name, description string, price float64, quantity int) {
	db := db.ConectDb()

	fmt.Println("Atualizando Produto: ", id, name, description, price, quantity)
	update, err := db.Prepare("update products set name=$1, description=$2, price=$3, quantity=$4 where id=$5")
	if err != nil {
		panic(err.Error())
	}

	fmt.Println("Atualizado")
	update.Exec(name, description, price, quantity, id)
	defer db.Close()
}
