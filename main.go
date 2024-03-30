package main

import (
	"astra/db"
	"astra/model"
	"encoding/json"
	"fmt"
	"path/filepath"

	"io/ioutil"
	"log"
	"net/http"
	"os"
)

func posthandler(w http.ResponseWriter, r *http.Request) {
	// Write response to the client
	fmt.Fprintf(w, "Hello, this is a simple HTTP server!")

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	emp := model.Data{}
	unerr := json.Unmarshal(body, &emp)
	if unerr != nil {
		log.Println("error in unmarshalling", unerr.Error())
		http.Error(w, unerr.Error(), http.StatusInternalServerError)
	}

	directory := "tmp/astra"
	if _, err := os.Stat(directory); os.IsNotExist(err) {
		if err := os.Mkdir(directory, 0755); err != nil {
			fmt.Println("Errorrrr:", err)
			return
		}
	}

	filePath := filepath.Join(directory, "data.txt")

	// Write the received data to a file
	if err := ioutil.WriteFile(filePath, body, 0644); err != nil {
		fmt.Println("Error:", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	/*insert data to db*/
	db.InsertDataToDB(emp)
}

func main() {
	http.HandleFunc("/postdata", posthandler)
	log.Println("server listening on port 8080")
	http.ListenAndServe(":8080", nil)
}
