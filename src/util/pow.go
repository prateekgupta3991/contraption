package util

import (
	"crypto/sha256"
	"encoding/hex"
	"strconv"
)

func CalculatePOW(prevBlockProof int64) int64 {
	proof := int64(1)
	for !ValidatePOW(prevBlockProof, proof) {
		proof++
	}
	return proof
}

func ValidatePOW(prevPow, pow int64) bool {
	prod := prevPow * pow
	prodB := []byte(strconv.Itoa(int(prod)))
	shaCode := sha256.New()
	shaCode.Write(prodB)
	sha := hex.EncodeToString(shaCode.Sum(nil))
	return sha[:4] == "0000"
}
