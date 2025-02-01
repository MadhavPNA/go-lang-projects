package main

import (
	"go-bookstore/pkg/routers"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func main() {
	r := mux.NewRouter()
	routers.RegisterBookStroeRoutes(r)
	http.Handle("/", r)
	log.Fatal(http.ListenAndServe(":9000", r))
}
