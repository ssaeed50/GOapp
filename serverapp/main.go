package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

func apihandler(w http.ResponseWriter, r *http.Request) {

	//fmt.Fprintf(w, "Hello World")
	log.Printf("Received %s request for host %s from IP address %s and X-FORWARDED-FOR %s",
		r.Method, r.Host, r.RemoteAddr, r.Header.Get("X-FORWARDED-FOR"))
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		body = []byte(fmt.Sprintf("error reading request body: %s", err))
	}
	resp := fmt.Sprintf("Hello, %s from Simple Server!", body)
	w.Write([]byte(resp))
	log.Printf("SimpleServer: Sent response %s", resp)
}
func healthzhandler(w http.ResponseWriter, r *http.Request) {

	fmt.Fprintf(w, "OK")
}

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
	mux := http.NewServeMux()
	//setuphandlers(mux)

	mux.HandleFunc("/healthz", healthzhandler)
	mux.HandleFunc("/api", apihandler)
	srv := &http.Server{
		ReadTimeout:       5 * time.Minute,
		WriteTimeout:      10 * time.Second,
		Addr:              "localhost:8080",
		IdleTimeout:       15 * time.Second,
		ReadHeaderTimeout: 3 * time.Second,
		Handler:           mux,
	}

	//log.Fatal(http.ListenAndServe(, mux))
	fmt.Printf("Starting server at port 8080\n")
	if err := srv.ListenAndServe(); err != nil {
		log.Fatal(err)
	}
}
