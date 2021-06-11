package services

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/praateekgupta3991/contraption/core"
)

type BlockchainService struct {
	bcn *core.Blockchain
}

func NewBlockchainService(b *core.Blockchain) *BlockchainService {
	return &BlockchainService{
		bcn: b,
	}
}

func (bcs *BlockchainService) GetChain(c *gin.Context) {
	chain, _ := bcs.bcn.GetChain()
	c.JSON(http.StatusOK, chain)
}
