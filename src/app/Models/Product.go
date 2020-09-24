package Models

type Product struct {
	Id       string `json: "id"`
	Name     string `json: "name"`
	Category string `json: "category"`
	Price    int    `json: "price"`
}
