package main

import (
	"bytes"
	"crypto/sha256"
	"strconv"
	"time"
)

// SetHash generates a hash of the block by concatenating all the data and hashing it
func (block *Block) SetHash() {
	// get the timestamp, convert digits to string vals (as bytes), e.g. [49 54 49 48 55 55 48 56 54 51]
	timestamp := []byte(strconv.FormatInt(block.Timestamp, 10))

	// concat all the data
	headers := bytes.Join([][]byte{timestamp, block.PreviousBlockHash, block.AllData}, []byte{})

	// hash the lot
	hash := sha256.Sum256(headers)

	// attach the hashed headers to the block
	block.MyBlockHash = hash[:]
}

// NewBlock creates a new block for the chain
func NewBlock(data string, prevBlockHash []byte) *Block {
	// create a new block
	block := &Block{
		Timestamp:         time.Now().Unix(),
		PreviousBlockHash: prevBlockHash,
		MyBlockHash:       []byte{},
		AllData:           []byte(data),
	}

	// hash the block and attach the hash to block.MyBlockHash
	block.SetHash()

	// return the new block
	return block
}

// NewGenesisBlock returns the genesis block - the first block in the chain
func NewGenesisBlock() *Block {
	return NewBlock("Genesis Block", []byte{})
}