package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
)

const (
	host     = "localhost"
	port     = 5443
	user     = "postgres"
	password = "555555"
	dbname   = "store"
)

func test(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Request URL: " + r.RequestURI)
	fmt.Println("Request method: " + r.Method)

	fmt.Println("Params: ")

	for k, v := range r.URL.Query() {
		fmt.Println(k + " = " + v[0])
	}
}

type Product struct {
	Id       string `json: "id"`
	Name     string `json: "name"`
	Category string `json: "category"`
	Price    int    `json: "price"`
}

var products = []Product{
	Product{Id: "1", Name: "Black Pan", Category: "for school", Price: 140},
	Product{Id: "2", Name: "Blue Pan", Category: "for school", Price: 150},
	Product{Id: "3", Name: "Red Pan", Category: "for school", Price: 140},
}

func main() {

	// postgres
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	err = db.Ping()
	if err != nil {
		panic(err)
	}

	fmt.Println("Successfully connected!")

	// router

	r := mux.NewRouter()

	r.HandleFunc("/products", getProducts).Methods("GET")
	r.HandleFunc("/product/{id}", getProduct).Methods("GET")
	r.HandleFunc("/product", createProduct).Methods("POST")
	r.HandleFunc("/product/{id}", updateProduct).Methods("PUT")
	r.HandleFunc("/product/{id}", deleteProduct).Methods("DELETE")

	// server
	fmt.Printf("Start server \n")

	if err := http.ListenAndServe(":8080", r); err != nil {
		log.Fatal(err)
	}

}

func getProducts(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		http.Error(w, "405 method not allowed", http.StatusMethodNotAllowed)
		return
	}
	w.Header().Set("Content-Type", "application/json")

	// postgres
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	err = db.Ping()
	if err != nil {
		panic(err)
	}

	rows, err := db.Query("select * from products")
	if err != nil {
		panic(err)
	}

	defer rows.Close()
	products := []Product{}

	for rows.Next() {
		p := Product{}
		err := rows.Scan(&p.Id, &p.Name, &p.Category, &p.Price)
		if err != nil {
			fmt.Println(err)
			continue
		}
		products = append(products, p)
	}

	for _, p := range products {
		fmt.Println(p.Id, p.Name, p.Category, p.Price)
	}

	json.NewEncoder(w).Encode(&products)

}

// Get Product by id
func getProduct(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		http.Error(w, "405 method not allowed", http.StatusMethodNotAllowed)
		return
	}

	w.Header().Set("Content-Type", "application/json")

	/*	params := mux.Vars(r)
		for _, item := range products {
			if item.Id == params["id"] {
				json.NewEncoder(w).Encode(item)
				return
			}
		} */

	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	err = db.Ping()
	if err != nil {
		panic(err)
	}

	params := mux.Vars(r)
	var ID = params["id"]

	row := db.QueryRow(`select * from products WHERE id = $1`, ID)

	products := []Product{}

	p := Product{}
	row.Scan(&p.Id, &p.Name, &p.Category, &p.Price)
	if err != nil {
		fmt.Println(err)
	}

	products = append(products, p)

	fmt.Println(p.Id, p.Name, p.Category, p.Price)
	//defer row.Close()

	json.NewEncoder(w).Encode(&products)

}

// Create Product item
func createProduct(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Error(w, "405 method not allowed", http.StatusMethodNotAllowed)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	var product Product
	_ = json.NewDecoder(r.Body).Decode(&product)
	product.Id = strconv.Itoa(rand.Intn(1000))
	products = append(products, product)
	json.NewEncoder(w).Encode(product)

}

func updateProduct(w http.ResponseWriter, r *http.Request) {
	if r.Method != "PUT" {
		http.Error(w, "405 method not allowed", http.StatusMethodNotAllowed)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for index, item := range products {
		if item.Id == params["id"] {
			products = append(products[:index], products[index+1:]...)
			var product Product
			_ = json.NewDecoder(r.Body).Decode(&product)
			product.Id = params["id"]
			products = append(products, product)
			json.NewEncoder(w).Encode(product)
			return
		}
	}
	json.NewEncoder(w).Encode(products)
}

func deleteProduct(w http.ResponseWriter, r *http.Request) {
	if r.Method != "DELETE" {
		http.Error(w, "405 method not allowed", http.StatusMethodNotAllowed)
		return
	}

	w.Header().Set("Content-Type", "application/json")

	/*
		params := mux.Vars(r)
		for index, item := range products {
			if item.Id == params["id"] {
				products = append(products[:index], products[index+1:]...)
				break
			}
		}
		json.NewEncoder(w).Encode(products)
	*/

}
