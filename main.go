package main

import (
	"fmt"
	"holyways/database"
	"holyways/pkg/mysql"
	"holyways/routes"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {

	// intial DB
	mysql.DatabaseInit()

	// run migration
	database.RunMigration()

	r := mux.NewRouter()

	routes.RouteInit(r.PathPrefix("/api/v1").Subrouter())

	fmt.Println("Server Running on Localhost: 5000")
	http.ListenAndServe("localhost:5000", r)

}
