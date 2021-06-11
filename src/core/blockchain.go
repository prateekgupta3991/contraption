package core

import (
	"time"

	"github.com/praateekgupta3991/contraption/entities"
)

type Blockchain struct {
	Genesis      *entities.Block
	CurrentBlock *entities.Block
	// NextBlock    *entities.Block
}

type BcnOperation interface {
	AddBlock(proof int64, prevHash string) (int64, error)
	GetChain() []string
}

func NewBlockchain() *Blockchain {
	bcn := new(Blockchain)
	bcn.Genesis = &entities.Block{
		Index:        1,
		Timestamp:    time.Now(),
		Proof:        100,
		PreviousHash: "calculate hash here",
	}
	bcn.CurrentBlock = bcn.Genesis
	return bcn
}

func (b *Blockchain) AddBlock(block *entities.Block) (int64, error) {
	b.CurrentBlock.NextBlock = block
	b.CurrentBlock = block
	return block.Index, nil
}

func (b *Blockchain) GetChain() ([]string, error) {
	chain := make([]string, 10)
	tmpBlk := b.Genesis
	for tmpBlk != nil {
		hashKey := tmpBlk.PreviousHash
		chain = append(chain, hashKey)
		tmpBlk = tmpBlk.NextBlock
	}
	return chain, nil
}
