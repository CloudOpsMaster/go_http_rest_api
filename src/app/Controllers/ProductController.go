package Controllers

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/Maks0123/go_http_rest_api/src/app/Models"
	"github.com/gorilla/mux"

)

const (
	host     = "localhost"
	port     = 5443
	user     = "postgres"
	password = "555555"
	dbname   = "store"
)

// Get all products

func GetProducts(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		http.Error(w, "405 method not allowed", http.StatusMethodNotAllowed)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	w.Header().Set("Access-Control-Allow-Origin", "*")

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
	products := []Models.Product{}

	for rows.Next() {
		p := Models.Product{}
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
func GetProduct(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		http.Error(w, "405 method not allowed", http.StatusMethodNotAllowed)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	w.Header().Set("Access-Control-Allow-Origin", "*")

	/*	params := mux.Vars(r)
		for _, item := range products {
			if item.Id == params["id"] {
				json.NewEncoder(w).Encode(item)
				return
			}
		} */
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

	params := mux.Vars(r)
	var ID = params["id"]

	row := db.QueryRow(`select * from products WHERE id = $1`, ID)

	products := []Models.Product{}

	p := Models.Product{}
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
func CreateProduct(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Error(w, "405 method not allowed", http.StatusMethodNotAllowed)
		return
	}

	w.Header().Set("Content-Type", "application/json")

	/*
		var product Product
		_ = json.NewDecoder(r.Body).Decode(&product)
		product.Id = strconv.Itoa(rand.Intn(1000))
		products = append(products, product)
		json.NewEncoder(w).Encode(product)

	*/

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
		fmt.Println(err)

	}

	var product Models.Product
	_ = json.NewDecoder(r.Body).Decode(&product)

	//id := product.Id
	//name := product.Name
	//category := product.Category
	//price := product.Price

	result, err := db.Exec(`INSERT INTO products(name,  category, price) VALUES($1, $2, $3)`, product.Name, product.Category, product.Price)
	if err != nil {
		panic(err)
	}

	fmt.Println("Product was created, id = " + product.Id)

	json.NewEncoder(w).Encode(&result)

}

func UpdateProduct(w http.ResponseWriter, r *http.Request) {
	if r.Method != "PUT" {
		http.Error(w, "405 method not allowed", http.StatusMethodNotAllowed)
		return
	}
	/* w.Header().Set("Content-Type", "application/json")
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

	*/
	w.Header().Set("Content-Type", "application/json")

	params := mux.Vars(r)

	var product Models.Product
	_ = json.NewDecoder(r.Body).Decode(&product)

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
	defer db.Close()

	var ID = params["id"]

	result, err := db.Exec(`UPDATE products SET name = $2, category = $3, price = $4 WHERE id = $1`, ID, product.Name, product.Category, product.Price)
	if err != nil {
		panic(err)
	}

	fmt.Println("Product was updated")

	json.NewEncoder(w).Encode(&result)

}

// Delete product by id
func DeleteProduct(w http.ResponseWriter, r *http.Request) {
	if r.Method != "DELETE" {
		http.Error(w, "405 method not allowed", http.StatusMethodNotAllowed)
		return
	}

	w.Header().Set("Content-Type", "application/json")

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
	defer db.Close()
	params := mux.Vars(r)

	var ID = params["id"]

	result, err := db.Exec(`DELETE FROM products WHERE id = $1`, ID)
	if err != nil {
		panic(err)
	}

	fmt.Println("Product was deleted")

	json.NewEncoder(w).Encode(&result)

}
