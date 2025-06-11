package main

import (
	"fmt"
	"log"
	"net/http"
	"service-api/routes"
	"service-api/utils/middleware"
	"service-api/database"
)

func main() {
	db := database.ConnectDB()
	database.Migrate(db);

	routers := routes.SetupRouters(db)
	routers.Use(middleware.Logger)
	
	addr := "localhost:8000"
	fmt.Printf("🚀 Server is running at http://%s\n", addr)

    if err := http.ListenAndServe(addr, routers); err != nil {
        log.Fatalf("Error starting server: %v", err)
    }
}