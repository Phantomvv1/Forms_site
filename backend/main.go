package main

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/jackc/pgx/v5"
)

type answers struct {
	Name                    string `json:"name"`
	Age                     string `json:"age"`
	RoutineChange           string `json:"routineChange"`
	NatureTime              string `json:"natureTime"`
	DailyHelp               bool   `json:"dailyHelp"`
	PublicTransport         bool   `json:"publicTransport"`
	Recycle                 string `json:"recycle"`
	FoodWaste               string `json:"foodWaste"`
	EcoProducts             string `json:"ecoProducts"`
	ClimateChangeImportance string `json:"climateChangeImportance"`
	PayMoreSustainable      bool   `json:"payMoreSustainable"`
}

type Response struct {
	Answers []answers `json:"answers"`
	Length  int       `json:"length"`
}

func getSubmittions(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

	if r.Method == http.MethodOptions {
		w.WriteHeader(http.StatusOK)
		return
	}

	conn, err := pgx.Connect(context.Background(), os.Getenv("DATABASE_URL"))
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close(context.Background())

	var response []answers

	rows, err := conn.Query(context.Background(), "select p.name, p.age, a.routine_change, a.nature_time, "+
		"a.help_the_environment, a.public_transport, a.recycle, "+
		"a.waste_food, a.use_eco_products, a.climate_change_importance, a.sustainable_energy "+
		"from people p "+
		"join answers a on p.submittion_id = a.submittion_id;")
	if err != nil {
		log.Fatal(err)
	}

	var length int

	for rows.Next() {
		var ans answers
		err = rows.Scan(
			&ans.Name,
			&ans.Age,
			&ans.RoutineChange,
			&ans.NatureTime,
			&ans.DailyHelp,
			&ans.PublicTransport,
			&ans.Recycle,
			&ans.FoodWaste,
			&ans.EcoProducts,
			&ans.ClimateChangeImportance,
			&ans.PayMoreSustainable,
		)

		if err != nil {
			log.Fatal(err)
		}

		response = append(response, ans)
		length += 1
	}

	var finalResponse Response = Response{
		Answers: response,
		Length:  length,
	}

	fmt.Println(finalResponse)

	w.Header().Set("Content-Type", "application/json")

	json.NewEncoder(w).Encode(finalResponse)
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
	defer conn.Close(context.Background())

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

	mux.Handle("/", http.FileServer(http.Dir(".")))

	mux.HandleFunc("/getSubmittions", getSubmittions)
	mux.HandleFunc("/sendData", recieveSubmittion)

	err := http.ListenAndServe(":42069", mux)
	if errors.Is(err, http.ErrServerClosed) {
		fmt.Println("Couldn't open server")
	} else if err != nil {
		log.Fatal(err)
	}
}
