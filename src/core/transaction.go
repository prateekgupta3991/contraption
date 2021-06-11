package core

import (
	"strconv"

	"github.com/praateekgupta3991/contraption/entities"
)

type TransationOperation interface {
}

func NewTransaction(bid int64, sndr, rcvr string, amount float64) entities.Transaction {
	strId := strconv.Itoa(int(bid))
	strTid := strId + "1" //fix this
	tid, _ := strconv.Atoi(strTid)
	return entities.Transaction{
		Id:       int64(tid),
		Sender:   sndr,
		Reciever: rcvr,
		Amount:   amount,
	}
}
