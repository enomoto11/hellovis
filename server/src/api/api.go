package api

import (
	"context"
	"fmt"
	"hellovis/ent"
	"hellovis/ent/migrate"
	"log"
	"os"

	"github.com/gin-gonic/gin"
)

type Controller interface {
	Register(r gin.IRouter)
}

func InitAPI() (*gin.Engine, *ent.Client) {
	router := gin.Default()

	path := fmt.Sprintf("%s:%s@tcp(db:3306)/%s?charset=utf8&parseTime=true", os.Getenv("MYSQL_USER"), os.Getenv("MYSQL_PASSWORD"), os.Getenv("MYSQL_DATABASE"))
	entClient, err := ent.Open("mysql", path)
	if err != nil {
		log.Fatalf("failed connect to mysql: %v", err)
	}

	ctx := context.Background()
	if err := entClient.Schema.Create(
		ctx,
		migrate.WithDropIndex(true),
		migrate.WithDropColumn(true),
	); err != nil {
		log.Fatalf("failed creating schema resources: %v", err)
	}

	// controllers := setUpControllers(entClient)
	// for _, controller := range controllers {
	// 	controller.Register(router)
	// }

	return router, entClient
}

// func setUpControllers(entClient *ent.Client) []Controller {

// 	return []Controller{}
// }
