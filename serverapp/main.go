package main

import (
	"database/sql"
	"fmt"
	"io"
	"log"
	"net/http"
	"strings"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

/// Upload function logic

func uploadfile(w http.ResponseWriter, r *http.Request) {

	log.Printf("Received %s request for host %s from IP address %s",
		r.Method, r.Host, r.RemoteAddr)

	/*file, err := os.Create("./result")

	if err != nil {

		panic(err)
	}
	*/
	// Return the number of bytes copied to newfile
	buf := new(strings.Builder)
	newfile, err := io.Copy(buf, r.Body)
	if err != nil {
		panic(err)
	}
	w.Write([]byte(fmt.Sprintf("%d bytes are recieved.\n", newfile)))
	fmt.Println("File content.\n", buf.String())

}

func connectdb(name string) {

	db, err := sql.Open("mysql", "root:VMware1!@tcp(localhost:3306)/test")

	if err != nil {
		panic(err.Error())
	}
	pingErr := db.Ping()
	if pingErr != nil {
		panic(pingErr)
	}
	fmt.Println("Connected!")

	//defer db.Close()
	/*_, err = db.Exec("CREATE DATABASE " + name)
	if err != nil {
		panic(err)
	}
	*/
	_, err = db.Exec("USE " + name)
	if err != nil {
		panic(err)
	}

	/*_, err = db.Exec("CREATE TABLE example ( id integer, data varchar(32) )")
	if err != nil {
		panic(err)
	}
	*/

}

func main() {

	mux := http.NewServeMux()
	// server struct
	srv := &http.Server{
		ReadTimeout:       5 * time.Minute,
		WriteTimeout:      10 * time.Second,
		Addr:              "localhost:8080",
		IdleTimeout:       15 * time.Second,
		ReadHeaderTimeout: 3 * time.Second,
		Handler:           mux,
	}

	connectdb("go")

	mux.HandleFunc("/upload", uploadfile)
	//log.Fatal(http.ListenAndServe(, mux))

	if err := srv.ListenAndServe(); err != nil {
		log.Fatal(err)
	}

}
