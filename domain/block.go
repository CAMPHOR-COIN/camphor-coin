package domain

import (
  "fmt"
  "time"
  "crypto/sha256"
)

type block struct {
	PrevHash  [32]byte
	Timestamp string
	Data      string
	Hash      [32]byte
}

// constructor
func NewBlock(prevHash [32]byte, timestamp string, data string, hash [32]byte) *block {
	return &block{
		PrevHash: prevHash,
		Timestamp: timestamp,
		Data: data,
		Hash: hash,
	}
}

// calculate hash
func CalculateHash(prevHash [32]byte, timestamp string, data string) [32]byte {
  sourceString := string(prevHash[:]) + timestamp + data
	return sha256.Sum256([]byte(sourceString))
}

// generate new block from transaction data
func GenerateNextBlock(data string, previousBlock block) *block {
	nextTimestamp := time.Now().String()
	nextHash := CalculateHash(previousBlock.Hash, nextTimestamp, data)
	return NewBlock(previousBlock.Hash, nextTimestamp, data, nextHash)
}

// generate very first block
func getGenesisBlock() *block {
	var genesisPrevHash, genesisHash [32]byte
	copy(genesisPrevHash[:], "0000000000000000000000000000000000000000000000000000000000000000")
	copy(genesisHash[:], "000000000019d6689c085ae165831e934ff763ae46a2a6c172b3f1b60a8ce26f")

	return NewBlock(genesisPrevHash, "1230974105", "4a5e1e4baab89f3a32518a88c31bc87f618f76673e2cc77ab2127b7afdeda33b", genesisHash)
}

// validate blocks
func IsValidNewBlock(newBlock, previousBlock *block) bool {
  if previousBlock.Hash != newBlock.PrevHash {
    fmt.Printf("invalid previous hash!")
    return false
  } else if CalculateHash(newBlock.PrevHash, newBlock.Timestamp, newBlock.Data) != newBlock.Hash {
    fmt.Printf("invalid hash!")
    return false
  }

  return true
}

func GetBlockchain() []*block{
	blockchain := []*block{getGenesisBlock()}
	return blockchain
}
