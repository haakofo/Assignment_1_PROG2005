package main

import (
	"assignment-1/handlers_and_structs"
	"log"
	"net/http"
	"os"
	"time"
)

func main() {
	handlers_and_structs.Elapsed_Time = time.Now()

	http.HandleFunc(handlers_and_structs.DEFAULT_TOTAL_PATH, handlers_and_structs.Default_Handler)
	http.HandleFunc(handlers_and_structs.NEIGHBOURUNIS_TOTAL_PATH, handlers_and_structs.Neighbour_unis_Handler)
	http.HandleFunc(handlers_and_structs.UNIINFO_PATH, handlers_and_structs.Uni_info_Handler)
	http.HandleFunc(handlers_and_structs.DIAG_PATH, handlers_and_structs.Diag_Handler)

	// Extract PORT variable from the environment variables - used in Heroku
	current_used_port := os.Getenv("PORT")

	// Override current_used_port with default current_used_port if not provided (e.g. local deployment)
	if current_used_port == "" {
		log.Println("$PORT has not been set. Default: 8080")
		current_used_port = "8080"
	}

	// Start HTTP server
	log.Println("Starting server on current_used_port " + current_used_port + " ...")
	log.Fatal(http.ListenAndServe(":"+current_used_port, nil))
}
