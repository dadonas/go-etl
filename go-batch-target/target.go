package main

import (
	"database/sql"
	"encoding/json"
	"io/ioutil"
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

	// Prepare statement for inserting data
	stmtIns, err := db.Prepare("INSERT INTO batch_target VALUES( ?, ?, ? )") // ? = placeholder
	if err != nil {
		panic(err.Error())
	}
	defer stmtIns.Close() // Close the statement when we leave main() / the program terminates

	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/target/records", func(w http.ResponseWriter, r *http.Request) {
		/*
			var batchs BatchRecords
			b, _ := ioutil.ReadAll(r.Body)
			json.Unmarshal(b, &batchs)

			for _, batch := range batchs {
				_, err = stmtIns.Exec(batch.ID, batch.Description, batch.Etc)
				if err != nil {
					panic(err.Error())
				}
			}
		*/

		var batch BatchRecord
		b, _ := ioutil.ReadAll(r.Body)
		json.Unmarshal(b, &batch)
		_, err = stmtIns.Exec(batch.ID, batch.Description, batch.Etc)
		if err != nil {
			panic(err.Error())
		}

	}).Methods("POST")

	log.Fatal(http.ListenAndServe(":8081", router))
}

//BatchRecord struct
type BatchRecord struct {
	ID          int    `json:"id"`
	Description string `json:"description"`
	Etc         string `json:"etc"`
}

//BatchRecords list
type BatchRecords []BatchRecord
