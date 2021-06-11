package main

import (
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"

	"github.com/praateekgupta3991/contraption/configs"
	"github.com/praateekgupta3991/contraption/core"
	"github.com/praateekgupta3991/contraption/services"
)

func main() {
	bcnMaster := core.NewBlockchain()
	blk := core.NewBlock(bcnMaster.CurrentBlock.Index, bcnMaster.Genesis.Proof, bcnMaster.CurrentBlock.PreviousHash)
	blk.Transactions = append(blk.Transactions, core.NewTransaction(bcnMaster.CurrentBlock.Index, "", "", 5))
	bcnMaster.AddBlock(blk)
	fmt.Println("The blockchain ran fine")

	var con *configs.Conf
	var err error
	if con, err = configs.InitConfig("./configs/conf.dev.json"); err != nil {
		log.Panicf("Error during config initialisation : %s", err.Error())
	}

	router := gin.New()
	router.Use(guidMiddleware())
	bcs := services.NewBlockchainService(bcnMaster)
	router.GET("/chain", bcs.GetChain)
	port := fmt.Sprintf(":%s", con.ServerPort)
	router.Run(port)
}

func guidMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		uuid := uuid.New()
		c.Set("uuid", uuid)
		fmt.Printf("The request with uuid %s is started \n", uuid)
		c.Next()
		fmt.Printf("The request with uuid %s is served \n", uuid)
	}
}
