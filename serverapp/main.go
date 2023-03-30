package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"time"
)

/// Upload function logic

func uploadfile(w http.ResponseWriter, r *http.Request) {

	log.Printf("Received %s request for host %s from IP address %s",
		r.Method, r.Host, r.RemoteAddr)

	//r.ParseMultipartForm(32 << 20)

	file, err := os.Create("./result")

	if err != nil {

		log.Fatal(err)
	}
	n, err := io.Copy(file, r.Body)
	if err != nil {
		panic(err)
	}
	w.Write([]byte(fmt.Sprintf("%d bytes are recieved.\n", n)))

}

//Test functions

/*func healthzhandler(w http.ResponseWriter, r *http.Request) {

	fmt.Fprintf(w, "OK")
}
*/
/*
func setuphandlers(mux *http.ServeMux) {

		mux.HandleFunc("/healthz", healthzhandler)
		mux.HandleFunc("/api", apihandler)
	}
*/

func main() {

	/*listenAddr := os.Getenv("LISTEN_ADDR")

	if len(listenAddr) == 0 {

		listenAddr = "3333"
	}
	*/
	//port := flag.String("port", "8080", "The http port, defaults to 8080")
	//setuphandlers(mux)
	//mux.HandleFunc("/healthz", healthzhandler)

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

	mux.HandleFunc("/upload", uploadfile)
	//log.Fatal(http.ListenAndServe(, mux))

	if err := srv.ListenAndServe(); err != nil {
		log.Fatal(err)
	}
}
