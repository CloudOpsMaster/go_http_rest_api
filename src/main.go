package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	// "os"

	"github.com/Maks0123/go_http_rest_api/src/app/Controllers"
	"github.com/Maks0123/go_http_rest_api/src/app/Models"
	"github.com/go-redis/redis"
	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
	//"github.com/joho/godotenv"

)

func test(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Request URL: " + r.RequestURI)
	fmt.Println("Request method: " + r.Method)

	fmt.Println("Params: ")

	for k, v := range r.URL.Query() {
		fmt.Println(k + " = " + v[0])
	}
}

/*
	// load .env file
	err := godotenv.Load()
	if err != nil {
	  log.Fata(err)
	}


 var cenvHost := os.Getenv("HOST")
 var envPort := os.Getenv( "PORT")
 var envUser := os.Getenv("USERNAME")
 var envPassword := os.Getenv("PASS")
 var envDBname := os.Getenv("NAME")



const (
	host     = envHost
	port     = envPort
	user     = envUser
	password = envPassword
	dbname   = envDBname
)

*/

const (
	host     = "localhost"
	port     = 5443
	user     = "postgres"
	password = "555555"
	dbname   = "store"
)

/*
type Product struct {
	Id       string `json: "id"`
	Name     string `json: "name"`
	Category string `json: "category"`
	Price    int    `json: "price"`
}
*/

var products = []Models.Product{
	Models.Product{Id: "1", Name: "Black Pan", Category: "for school", Price: 140},
	Models.Product{Id: "2", Name: "Blue Pan", Category: "for school", Price: 150},
	Models.Product{Id: "3", Name: "Red Pan", Category: "for school", Price: 140},
}

type JsonResponse struct {
	Type    string           `json:"type"`
	Data    []Models.Product `json:"data"`
	Message string           `json:"message"`
}

func main() {

	// Redis
	client := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})

	json, err := json.Marshal(Models.Product{Id: "12345", Name: "Elliot", Category: "Game", Price: 25})
	if err != nil {
		fmt.Println(err)
	}

	err = client.Set("id1234", json, 0).Err()
	if err != nil {
		fmt.Println(err)
	}
	val, err := client.Get("id1234").Result()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(val)

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

	r.HandleFunc("/products", Controllers.GetProducts).Methods("GET")
	r.HandleFunc("/product/{id}", Controllers.GetProduct).Methods("GET")
	r.HandleFunc("/product", Controllers.CreateProduct).Methods("POST")
	r.HandleFunc("/product/{id}", Controllers.UpdateProduct).Methods("PUT")
	r.HandleFunc("/product/{id}", Controllers.DeleteProduct).Methods("DELETE")

	// server
	fmt.Printf("Start server \n")

	if err := http.ListenAndServe(":8080", r); err != nil {
		log.Fatal(err)
	}

}
