package controllers

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"strconv"

	"github.com/alexsantossilva/go-app/models"
)

var templates = template.Must(template.ParseGlob("templates/*.html"))

func Index(w http.ResponseWriter, r *http.Request) {
	products := models.FindAll()
	templates.ExecuteTemplate(w, "Index", products)
}

func New(w http.ResponseWriter, r *http.Request) {
	templates.ExecuteTemplate(w, "New", nil)
}

func Insert(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		name := r.FormValue("name")
		description := r.FormValue("description")
		price := r.FormValue("price")
		quantity := r.FormValue("quantity")

		priceConvert, err := strconv.ParseFloat(price, 64)
		if err != nil {
			log.Println("Erro na conversação de preço: ", err)
		}

		quantityConvert, err := strconv.Atoi(quantity)
		if err != nil {
			log.Println("Erro na conversação de preço: ", err)
		}

		models.InsertProduct(name, description, priceConvert, quantityConvert)
	}
	http.Redirect(w, r, "/", 301)
}

func Delete(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")
	fmt.Println("ID", id)

	models.Delete(id)
	http.Redirect(w, r, "/", 301)
}

func Edit(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")
	product := models.FindById(id)
	templates.ExecuteTemplate(w, "Edit", product)
}

func Update(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Method: ", r.Method)
	if r.Method == "POST" {
		id := r.FormValue("id")
		name := r.FormValue("name")
		description := r.FormValue("description")
		price := r.FormValue("price")
		quantity := r.FormValue("quantity")

		idConvert, err := strconv.Atoi(id)
		if err != nil {
			log.Println("Erro na conversação do id: ", err)
		}

		priceConvert, err := strconv.ParseFloat(price, 64)
		if err != nil {
			log.Println("Erro na conversação de preço: ", err)
		}

		quantityConvert, err := strconv.Atoi(quantity)
		if err != nil {
			log.Println("Erro na conversação de preço: ", err)
		}

		models.Update(idConvert, name, description, priceConvert, quantityConvert)
	}

	http.Redirect(w, r, "/", 301)
}
