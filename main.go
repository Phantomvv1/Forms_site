package main

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"

	"github.com/jackc/pgx/v5"
)

type answers struct {
	Name                    string
	Age                     string
	RoutineChange           string
	NatureTime              string
	DailyHelp               bool
	PublicTransport         bool
	Recycle                 string
	FoodWaste               string
	EcoProducts             string
	ClimateChangeImportance string
	PayMoreSustainable      bool
}

func defaultFunc(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

	if r.Method == http.MethodOptions {
		w.WriteHeader(http.StatusOK)
		return
	}

	fmt.Println("Hello")
	io.WriteString(w, "default")
}

func getSubmittions(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

	if r.Method == http.MethodOptions {
		w.WriteHeader(http.StatusOK)
		return
	}

}

func recieveSubmittion(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

	if r.Method == http.MethodOptions {
		w.WriteHeader(http.StatusOK)
		return
	}

	var rawData answers
	err := json.NewDecoder(r.Body).Decode(&rawData)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(rawData)

	conn, err := pgx.Connect(context.Background(), os.Getenv("DATABASE_URL"))
	if err != nil {
		log.Fatal(err)
	}

	_, err = conn.Exec(context.Background(), "create table if not exists people (submittion_id serial primary key, name text, age int);")
	if err != nil {
		log.Fatal(err)
	}
	_, err = conn.Exec(context.Background(), "insert into people (name, age) values ($1, $2);", rawData.Name, rawData.Age)
	if err != nil {
		log.Fatal(err)
	}

	_, err = conn.Exec(context.Background(), "create table if not exists answers (submittion_id serial primary key, "+
		"routine_change text, nature_time int, help_the_environment bool, "+
		"public_transport bool, recycle text, waste_food text, "+
		"use_eco_products text, climate_change_importance text, "+
		"sustainable_energy bool);")
	if err != nil {
		log.Fatal(err)
	}

	_, err = conn.Exec(context.Background(), "insert into answers  (routine_change, nature_time, "+
		"help_the_environment, public_transport, recycle, waste_food, "+
		"use_eco_products, climate_change_importance, "+
		"sustainable_energy) values ($1, $2, $3, $4, $5, $6, $7, $8, $9);", rawData.RoutineChange, rawData.NatureTime,
		rawData.DailyHelp, rawData.PublicTransport, rawData.Recycle, rawData.FoodWaste,
		rawData.EcoProducts, rawData.ClimateChangeImportance, rawData.PayMoreSustainable)
	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("/getSubmittions", getSubmittions)
	mux.HandleFunc("/sendData", recieveSubmittion)
	mux.HandleFunc("/", defaultFunc)

	err := http.ListenAndServe(":42069", mux)
	if errors.Is(err, http.ErrServerClosed) {
		fmt.Println("Couldn't open server")
	} else if err != nil {
		log.Fatal(err)
	}
}
