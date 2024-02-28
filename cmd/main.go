package main

import (
	"bookstore-go/internals/routes"
	"bookstore-go/pkg"
	"log"
)

// Depedency Injection (DI)

func main() {
	// inisialisasi DB
	_, err := pkg.InitMySQL()
	if err != nil {
		log.Fatal(err)
		// return
	}

	// inisialisasi Router
	router := routes.InitRouter()
	// inisialisasi Server
	server := pkg.InitServer(router)

	// Jalankan server
	if err := server.ListenAndServe(); err != nil {
		log.Fatal(err)
	}
}
