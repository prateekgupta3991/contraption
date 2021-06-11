package core

import (
	"fmt"
	"strconv"

	"github.com/praateekgupta3991/contraption/entities"
)

type TransactionService struct {
}

type TransactionOperation interface {
	CreateTransaction(bid int64, sndr, rcvr string, amount float64) entities.Transaction
	GetTransactionId(bid int64) int64
}

func NewTransaction(bid int64, sndr, rcvr string, amount float64) entities.Transaction {
	strId := strconv.Itoa(int(bid))
	strTid := fmt.Sprintf("%s%s", strId, "1") //fix this
	tid, _ := strconv.Atoi(strTid)
	return entities.Transaction{
		Id:       int64(tid),
		Sender:   sndr,
		Reciever: rcvr,
		Amount:   amount,
	}
}

func (t *TransactionService) CreateTransaction(bid int64, sndr, rcvr string, amount float64) entities.Transaction {
	strId := strconv.Itoa(int(bid))
	strTid := fmt.Sprintf("%s%s", strId, "1")
	tid, _ := strconv.Atoi(strTid)
	return entities.Transaction{
		Id:       int64(tid),
		Sender:   sndr,
		Reciever: rcvr,
		Amount:   amount,
	}
}

func (t *TransactionService) GetTransactionId(bid int64) int64 {
	strId := strconv.Itoa(int(bid))
	strTid := fmt.Sprintf("%s%s", strId, "1")
	tid, _ := strconv.Atoi(strTid)
	return int64(tid)
}
