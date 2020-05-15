package main

import (
	"github.com/BalusChen/IKHNAIE_API/client"
	"github.com/BalusChen/IKHNAIE_API/views"
	"github.com/gin-gonic/gin"
)

func main() {
	helloCC()

	router := gin.Default()

	views.InitRoutes(router)

	panic(router.Run(":9877"))
}

const (
	org1CfgPath = "./config/org1sdk-config.yaml"
	org2CfgPath = "./config/org2sdk-config.yaml"

	peer0Org1 = "peer0.org1.example.com"
	peer0Org2 = "peer0.org2.example.com"
)

func helloCC() {
	org1Client := client.New(org1CfgPath, "Org1", "Admin", "User1")
	if err := org1Client.Query(peer0Org1, "a"); err != nil {
		panic(err)
	}
}
