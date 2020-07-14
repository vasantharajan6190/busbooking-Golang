package main

import (
	"log"
	"net/http"
    "routes"
	"github.com/gorilla/mux"
)

//main function
func main() {
	r := mux.NewRouter().StrictSlash(true)
	//users
	r.HandleFunc("/users/{name}", routes.getusers).Methods("GET", "OPTIONS")
	r.HandleFunc("/users", routes.addusers).Methods("POST", "OPTIONS")
	//travels
	r.HandleFunc("/travels", routes.gettravels).Methods("GET", "OPTIONS")
	r.HandleFunc("/travels/{name}", routes.gettravelsbyname).Methods("GET", "OPTIONS")
	r.HandleFunc("/travels", routes.addtravels).Methods("POST", "OPTIONS")
	//bookings
	r.HandleFunc("/bookings", routes.getallbookings).Methods("GET", "OPTIONS")
	r.HandleFunc("/bookings/{name}", routes.getbookings).Methods("GET", "OPTIONS")
	r.HandleFunc("/bookings", routes.addbookings).Methods("POST", "OPTIONS")
	r.HandleFunc("/bookings", routes.updatebookings).Methods("PUT", "OPTIONS")
	log.Fatal(http.ListenAndServe(":8000", r))
}
