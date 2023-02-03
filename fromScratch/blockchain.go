package main

// AddBlock adds a new block to the blockchain
func (blockchain *Blockchain) AddBlock(data string) {
	// select the previous block
	previousBlock := blockchain.Blocks[len(blockchain.Blocks)-1]

	// create the new block and pass in the data
	newBlock := NewBlock(data, previousBlock.MyBlockHash)

	// append the new block to the chain
	blockchain.Blocks = append(blockchain.Blocks, newBlock)
}

// NewBlockchain creates a new blockchain and adds the genesis block to it.
func NewBlockchain() *Blockchain {
	return &Blockchain{[]*Block{NewGenesisBlock()}}
}