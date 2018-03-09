package main

import (
	"bytes"
	"crypto/sha256"
	"fmt"
	"strconv"
	"time"
)

// Blockchain is our global blockchain.
var Blockchain []Block

// Block is our basic data structure!
type Block struct {
	Data      string
	Timestamp int64
	PrevHash  []byte
	Hash      []byte
}

// InitBlockchain creates our first Genesis node.
func InitBlockchain() {
	genesisBlock  := Block{"genesisBlock", time.Now().Unix(), []byte{}, []byte{}}
	genesisBlock.Hash = genesisBlock.calculateHash()
	Blockchain = []Block{genesisBlock}
}

// NewBlock creates a new Blockchain Block.
func NewBlock(oldBlock Block, data string) Block {
	block  := Block{data, time.Now().Unix(), []byte{}, []byte{}}
	block.PrevHash = oldBlock.Hash
	block.Hash = block.calculateHash()
	return block
}

// AddBlock adds a new block to the Blockchain.
func AddBlock(b Block) error {
	lastBlock := Blockchain[len(Blockchain)-1]
	if bytes.Equal(b.PrevHash, lastBlock.Hash) {
		Blockchain = append(Blockchain, b)
		return nil
	}
	return fmt.Errorf("New block contains invalid previous hash string")
}

func (b *Block) calculateHash() []byte {
	timestamp := []byte(strconv.FormatInt(b.Timestamp, 10))
	data := []byte(b.Data)
	headers := bytes.Join([][]byte{b.PrevHash, data, timestamp}, []byte{})
	hash := sha256.Sum256(headers)
	return hash[:]
}
