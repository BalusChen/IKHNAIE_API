package main

import (
	"github.com/BalusChen/IKHNAIE_API/views"
	"github.com/gin-contrib/pprof"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	pprof.Register(router)
	views.InitRoutes(router)

	panic(router.Run(":9877"))
}
