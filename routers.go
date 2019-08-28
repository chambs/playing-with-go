package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// Product is a product class
type Product struct {
	ID    int
	Name  string
	Price float32
}

// HandleFoo handles foo
func HandleFoo(res http.ResponseWriter, req *http.Request) {
	fmt.Fprint(res, "This is foo")
}

// HandleJSON handles json
func HandleJSON(res http.ResponseWriter, req *http.Request) {
	jsonMap := make(map[string]string)
	jsonMap["Id"] = "123"
	toJSON, err := json.Marshal(jsonMap)

	if err != nil {
		fmt.Println("error marshal", err)
	} else {
		// fmt.Println("marshal ok", string(toJSON))
	}
	fmt.Fprint(res, string(toJSON))
}

// HandleProduct handles /product
func HandleProduct(res http.ResponseWriter, req *http.Request) {
	products := make([]Product, 4, 4)

	products[0] = Product{1, "Pen", 1.0}
	products[1] = Product{2, "Chocolate", 2.50}
	products[2] = Product{3, "Book", 10.0}
	products[3] = Product{4, "Fan", 999.99}

	products = append(products, Product{5, "T-Shirt", 9.75})
	// fmt.Printf("arrray size: %d\n", len(products))

	productsAsJSON, err := json.Marshal(products)

	if err != nil {
		fmt.Println("Error", err)
	} else {
		// fmt.Println(string(productsAsJSON))
		res.Header().Set("Content-Type", "application/json")
		fmt.Fprint(res, string(productsAsJSON))
	}
}
