package entities

import "time"

type Block struct {
	Index        int64         `json:"index"`
	Timestamp    time.Time     `json:"timestamp"`
	Transactions []Transaction `json:"transations"`
	Proof        int64         `json:"proof"`
	PreviousHash string        `json:"previousHash"`
	NextBlock    *Block        `json:"next"`
}

type Transaction struct {
	Id       int64   `json:"id"`
	Sender   string  `json:"sender"`
	Reciever string  `json:"reciever"`
	Amount   float64 `json:"amount"`
}
