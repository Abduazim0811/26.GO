package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

func main() {
	dsn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", "localhost", 5432, "postgres", "Abdu0811", "users")
	db, err := sql.Open("postgres", dsn)
	if err != nil {
		fmt.Println("error")
		log.Fatal(err)
	}
	defer db.Close()
	query := `
SELECT p.product_name, p.unit, p.price, c.category_name, c.description
FROM products AS p
JOIN categories AS c ON p.category_id = c.category_id
WHERE c.category_name = 'Beverages';
`

	rows, err := db.Query(query)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	for rows.Next() {
		var productName, unit, description, categoryName string
		var price float64
		if err := rows.Scan(&productName, &unit, &price, &categoryName, &description); err != nil {
			log.Fatal(err)
		}
		fmt.Printf("Product Name: %s\nUnit: %s\nPrice: %.2f\nCategory: %s\nDescription: %s\n\n",
			productName, unit, price, categoryName, description)
	}

	if err := rows.Err(); err != nil {
		log.Fatal(err)
	}
}
