package services

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/praateekgupta3991/contraption/core"
	"github.com/praateekgupta3991/contraption/entities"
)

type BlockchainService struct {
	bcn core.BcnOperation
	blk core.BlockOperation
	txn core.TransactionOperation
}

func NewBlockchainService(bop core.BcnOperation, block core.BlockOperation, transaction core.TransactionOperation) *BlockchainService {
	return &BlockchainService{
		bcn: bop,
		blk: block,
		txn: transaction,
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
			nBlock := bcs.blk.CreateBlock(bcs.bcn.GetIndex(), bcs.bcn.GetProof(), bcs.bcn.GetPrevHash())
			txnEntity.Id = bcs.txn.GetTransactionId(bcs.bcn.GetIndex())
			nBlock.Transactions = []entities.Transaction{*txnEntity}
			bid, _ := bcs.bcn.AddBlock(nBlock)
			c.JSON(http.StatusOK, bid)
		}
	}
}
