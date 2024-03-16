package main

import (
	_ "github.com/lib/pq"
	"github.com/sawitpro/technical_test/handler"
)

func main() {
	// err := godotenv.Load("./config/.env")
	// if err != nil {
	// 	panic(".env is not loaded properly")
	// }

	server := handler.InitServer()
	server.Run()
}
