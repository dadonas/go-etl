package main

import (
	"database/sql"
	"encoding/json"
	"log"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
)

func main() {
	db, err := sql.Open("mysql", "root:root@/go_batch")

	if err != nil {
		panic(err.Error())
	}

	defer db.Close()

	rows, err := db.Query("SELECT * FROM batch_source")

	batchs := make([]BatchRecord, 0)

	for rows.Next() {
		var batch BatchRecord
		err = rows.Scan(&batch.ID, &batch.Description, &batch.Etc)
		batchs = append(batchs, batch)
	}

	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/source/records", func(w http.ResponseWriter, r *http.Request) {
		json.NewEncoder(w).Encode(batchs)
	})

	log.Fatal(http.ListenAndServe(":8080", router))
}

//BatchRecord struct
type BatchRecord struct {
	ID          int    `json:"id"`
	Description string `json:"description"`
	Etc         string `json:"etc"`
}

//BatchRecords list
type BatchRecords []BatchRecord
