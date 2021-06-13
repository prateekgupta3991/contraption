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
	GetChain() ([]entities.Block, error)
	GetIndex() int64
	GetProof() int64
	GetPrevHash() string
	GetCurrentBlock() *entities.Block
	UpdateChain(*entities.Block)
}

func NewBlockchain() *Blockchain {
	bcn := new(Blockchain)
	bcn.Genesis = &entities.Block{
		Index:        1,
		Timestamp:    time.Now(),
		Proof:        100,
		PreviousHash: "1",
	}
	bcn.CurrentBlock = bcn.Genesis
	return bcn
}

func (b *Blockchain) AddBlock(block *entities.Block) (int64, error) {
	b.CurrentBlock.NextBlock = block
	b.CurrentBlock = block
	return block.Index, nil
}

func (b *Blockchain) GetChain() ([]entities.Block, error) {
	chain := make([]entities.Block, 0)
	tmpBlk := b.Genesis
	for tmpBlk != nil {
		chain = append(chain, *tmpBlk)
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

func (b *Blockchain) GetCurrentBlock() *entities.Block {
	return b.CurrentBlock
}

func (b *Blockchain) UpdateChain(syncBlk *entities.Block) {
	b.Genesis = syncBlk
}
