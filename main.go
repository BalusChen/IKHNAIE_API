package main

import (
	"github.com/BalusChen/IKHNAIE_API/views"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	views.InitRoutes(router)

	panic(router.Run(":9877"))
}
