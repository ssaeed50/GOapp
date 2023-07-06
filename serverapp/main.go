package main

import (
	"crypto/tls"
	"database/sql"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
)

/// Upload Handler logic

func uploadfile(w http.ResponseWriter, r *http.Request) {

	log.Printf("Received %s request for host %s from IP address %s",
		r.Method, r.Host, r.RemoteAddr)

	// Get the file content and name from the request

	fileContent, _, err := r.FormFile("file")

	if err != nil {

		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	defer fileContent.Close()

	// Read the file content
	fileBytes, err := ioutil.ReadAll(fileContent)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	//fileName := fileHeader.Filename

	//fmt.Println("Filename: %s\n",fileName)
	//fmt.Println("content: %s\n",fileContent)

	// Connect to the MySQL database
	db, err := sql.Open("mysql", "root:VMware1!@tcp(localhost:3306)/test")

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	pingErr := db.Ping()
	if pingErr != nil {
		http.Error(w, pingErr.Error(), http.StatusInternalServerError)
		return
	}
	defer db.Close()

	// Insert the file content and name into the database

	result, err := db.Exec("INSERT INTO files (content) VALUES (?)", fileBytes)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	log.Printf("Saved file to the database. Rows affected: %d\n", rowsAffected)

	// Return a success response to the client
	fmt.Fprint(w, "File uploaded successfully")

}

func main() {

	//Create a TLS listner

	cert, err := tls.LoadX509KeyPair("server.crt", "server.key")
	if err != nil {
		log.Fatalf("Failed to load TLS certificate: %v", err)
	}

	config := &tls.Config{Certificates: []tls.Certificate{cert}}
	listener, err := tls.Listen("tcp", ":8080", config)
	if err != nil {
		log.Fatalf("Failed to create listener: %v", err)
	}
	defer listener.Close()

	// Handle incoming requests
	http.HandleFunc("/upload", uploadfile)
	log.Printf("Server listening on %s", listener.Addr().String())
	http.Serve(listener, nil)

}
