package main
import (
	"fmt"
	"log"
	"net/http"
	"service-api/routes"
	"service-api/utils/middleware"
)

func main() {
	routers := routes.SetupRouters()
	routers.Use(middleware.Logger)
	
	addr := "localhost:8000"
	fmt.Printf("🚀 Server is running at http://%s\n", addr)

    if err := http.ListenAndServe(addr, routers); err != nil {
        log.Fatalf("Error starting server: %v", err)
    }
}