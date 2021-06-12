package main

import (
	"fmt"
	"log"
	"net"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"

	"github.com/praateekgupta3991/contraption/clients"
	"github.com/praateekgupta3991/contraption/configs"
	"github.com/praateekgupta3991/contraption/core"
	"github.com/praateekgupta3991/contraption/services"
	"github.com/praateekgupta3991/contraption/util"
)

func main() {
	bcnMaster := core.NewBlockchain()
	byteOfStruct := []byte(fmt.Sprintf("%v", bcnMaster.CurrentBlock))
	blockHashVal := util.GetShaHash(byteOfStruct)
	blk := core.NewBlock(bcnMaster.CurrentBlock.Index, bcnMaster.Genesis.Proof, blockHashVal)
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
	port := fmt.Sprintf(":%s", con.ServerPort)

	blkSer := &core.BlockService{}
	trxnSer := &core.TransactionService{}
	hClient := GetHttpClient()
	inClient := clients.InitInterNodeClient(hClient)
	bcs := services.NewBlockchainService(bcnMaster, blkSer, trxnSer, getIp(port), inClient)
	router.GET("/chain", bcs.GetChain)
	router.POST("/txn", bcs.NewTxn)
	router.POST("/nodes/register", bcs.RegisterNewNodes)
	router.POST("/nodes/resolve", bcs.ResolveChain)

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

func getIp(port string) string {
	var myIp string
	ifaces, _ := net.Interfaces()
	for _, i := range ifaces {
		addrs, _ := i.Addrs()
		for _, addr := range addrs {
			switch v := addr.(type) {
			case *net.IPNet:
				myIp = v.IP.String()
			case *net.IPAddr:
				myIp = v.IP.String()
			}
		}
	}
	return fmt.Sprintf("%s:%s", myIp, port)
}

func GetHttpClient() *http.Client {
	tr := &http.Transport{
		MaxIdleConns:       10,
		IdleConnTimeout:    30 * time.Second,
		DisableCompression: true,
	}
	return &http.Client{
		Transport: tr,
	}
}
