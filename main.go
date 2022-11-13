package main

import (
	"github.com/gin-gonic/gin"
	app "github.com/vandenbill/brand-commerce/App"
)

func main() {
	r := gin.Default()
	app.NewRouter(r)
	r.Run()
}
