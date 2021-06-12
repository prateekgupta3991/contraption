package services

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/praateekgupta3991/contraption/clients"
	"github.com/praateekgupta3991/contraption/core"
	"github.com/praateekgupta3991/contraption/entities"
	"github.com/praateekgupta3991/contraption/util"
)

type BlockchainService struct {
	bcn      core.BcnOperation
	blk      core.BlockOperation
	txn      core.TransactionOperation
	nodes    []string
	inClient clients.InterNodeComm
}

func NewBlockchainService(bop core.BcnOperation, block core.BlockOperation, transaction core.TransactionOperation, ip string, incomm clients.InterNodeComm) *BlockchainService {
	return &BlockchainService{
		bcn:      bop,
		blk:      block,
		txn:      transaction,
		nodes:    []string{ip},
		inClient: incomm,
	}
}

func (bcs *BlockchainService) GetChain(c *gin.Context) {
	chain, _ := bcs.bcn.GetChain()
	c.JSON(http.StatusOK, chain)
}

func (bcs *BlockchainService) NewTxn(c *gin.Context) {
	if body, err := ioutil.ReadAll(c.Request.Body); err != nil {
		fmt.Printf("Error encountered : %v", err.Error())
		c.JSON(http.StatusBadRequest, "Bad request")
	} else {
		txnEntity := new(entities.Transaction)
		err := json.Unmarshal(body, &txnEntity)
		if err != nil {
			fmt.Printf("Could not process the webhook. Error encountered : %v\n", err.Error())
			c.JSON(http.StatusBadRequest, "Bad request")
		} else {
			byteOfStruct := []byte(fmt.Sprintf("%v", bcs.bcn.GetCurrentBlock()))
			blockHashVal := util.GetShaHash(byteOfStruct)
			nBlock := bcs.blk.CreateBlock(bcs.bcn.GetIndex(), bcs.bcn.GetProof(), blockHashVal)
			txnEntity.Id = bcs.txn.GetTransactionId(bcs.bcn.GetIndex())
			nBlock.Transactions = []entities.Transaction{*txnEntity}
			bid, _ := bcs.bcn.AddBlock(nBlock)
			c.JSON(http.StatusOK, bid)
		}
	}
}

func (bcs *BlockchainService) RegisterNewNodes(c *gin.Context) {
	isPresent := false
	qp := c.Request.URL.Query()
	ip := qp.Get("ip")
	for _, v := range bcs.nodes {
		if v == ip {
			isPresent = true
		}
	}
	if !isPresent {
		bcs.nodes = append(bcs.nodes, ip)
		c.JSON(http.StatusOK, gin.H{"info": "Node added"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"info": "Node already present"})
}

func (bcs *BlockchainService) ResolveChain(c *gin.Context) {
	for i, v := range bcs.nodes {
		if i == 0 {
			continue
		}
		chain, err := bcs.inClient.Chain(v)
		if err != nil {
			c.JSON(http.StatusOK, gin.H{"error": "Failed to get the chain"})
			return
		}
		fmt.Println("Printing the chain")
		for _, val := range chain {
			fmt.Println(val)
		}
	}
	c.JSON(http.StatusOK, gin.H{"info": "Got the chain"})
}
