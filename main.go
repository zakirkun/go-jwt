package main

import (
	"github.com/gin-gonic/gin"
	app "github.com/vandenbill/brand-commerce/App"
	"log"
)

func main() {
	app.Logger()

	db := app.NewDB()

	r := gin.Default()
	app.L.Info("Load gin engine")

	app.NewRouter(r, db)
	app.L.Info("Load router")

	err := r.Run()
	if err != nil {
		log.Panic(err)
	}
}
