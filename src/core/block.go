package core

import (
	"fmt"
	"time"

	"github.com/praateekgupta3991/contraption/entities"
	"github.com/praateekgupta3991/contraption/util"
)

type BlockOperation interface {
}

func NewBlock(prevBid, prevProof int64, prevHash string) *entities.Block {
	blk := &entities.Block{
		Index:        prevBid + 1,
		Timestamp:    time.Now(),
		Proof:        util.CalculatePOW(prevProof),
		PreviousHash: prevHash,
	}
	fmt.Printf("Proof for the block %d - %d", blk.Index, blk.Proof)
	return blk
}
