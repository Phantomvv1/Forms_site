package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/jackc/pgx/v5"
)

func main() {
	conn, err := pgx.Connect(context.Background(), os.Getenv("DATABASE_URL"))
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close(context.Background())

	var id int
	var name, contactName, addres, city, postalCode, country string
	err = conn.QueryRow(context.Background(), "select * from customers where customer_id=1").Scan(&id, &name, &contactName, &addres, &city, &postalCode, &country)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("ID: %v, Name: %v, Contact name: %v, Adress: %v, City: %v, Postal code: %v, Country: %v",
		id, name, contactName, addres, city, postalCode, country)
}
