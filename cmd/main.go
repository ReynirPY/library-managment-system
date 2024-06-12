package main

import (
	"net/http"

	"github.com/ReynirPY/library-managment-system/api"
	"github.com/ReynirPY/library-managment-system/config"
)

func main() {
	config.InitDB()

	http.ListenAndServe(":8080", api.RegisterRoutes())
}
