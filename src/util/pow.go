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
	shaVal := GetShaHash(prodB)
	return shaVal[:4] == "0000"
}

func GetShaHash(val []byte) string {
	shaCode := sha256.New()
	shaCode.Write(val)
	sha := hex.EncodeToString(shaCode.Sum(nil))
	return sha
}
