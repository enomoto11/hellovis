package main

import (
	"hellovis/api"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	router, entClient := api.InitAPI()
	defer entClient.Close()

	router.Run()
}
