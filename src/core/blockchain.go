package core

import (
	"time"

	"github.com/praateekgupta3991/contraption/entities"
)

type Blockchain struct {
	Genesis      *entities.Block
	CurrentBlock *entities.Block
}

type BcnOperation interface {
	AddBlock(block *entities.Block) (int64, error)
	GetChain() ([]string, error)
	GetIndex() int64
	GetProof() int64
	GetPrevHash() string
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

func (b *Blockchain) GetIndex() int64 {
	return b.CurrentBlock.Index
}

func (b *Blockchain) GetProof() int64 {
	return b.CurrentBlock.Proof
}

func (b *Blockchain) GetPrevHash() string {
	return b.CurrentBlock.PreviousHash
}
