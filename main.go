package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"time"
)

func main() {
	m := http.NewServeMux()

	m.HandleFunc("/", handlePage)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080" // default to port 8080
	}
	srv := http.Server{
		Handler:      m,
		Addr:         ":" + port, // construct the address
		WriteTimeout: 30 * time.Second,
		ReadTimeout:  30 * time.Second,
	}

	// this blocks forever, until the server
	// has an unrecoverable error
	//fmt.Println("server started on ", port)
	fmt.Println("Server will start on ", srv.Addr)
	err := srv.ListenAndServe()
	log.Fatal(err)
}

func handlePage(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	w.WriteHeader(200)
	const page = `<html>
<head></head>
<body>
	<p> Hi Docker, I pushed a new version! </p>
</body>
</html>
`
	w.Write([]byte(page))
}
