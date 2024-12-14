package main

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"strconv"
	"strings"
)

type Block struct {
	Data     int
	PrevHash string
	Nonce    int
	Hash     string
}

type Blockchain struct {
	Blocks     []Block
	Difficulty int
}

func hash(data int, prevHash string, nonce int) string {
	input := strconv.Itoa(data) + prevHash + strconv.Itoa(nonce)
	hash := sha256.Sum256([]byte(input))
	return hex.EncodeToString(hash[:])
}

// Performs proof-of-work by finding a valid nonce for the block.
func mineBlock(data int, prevHash string, difficulty int) Block {
	prefix := strings.Repeat("0", difficulty)
	var nonce int
	var hashValue string

	for {
		hashValue = hash(data, prevHash, nonce)
		if strings.HasPrefix(hashValue, prefix) {
			break
		}
		nonce++
	}

	return Block{
		Data:     data,
		PrevHash: prevHash,
		Nonce:    nonce,
		Hash:     hashValue,
	}
}

func (bc *Blockchain) addBlock(data int) {
	var prevHash string
	if len(bc.Blocks) == 0 {
		prevHash = ""
	} else {
		prevHash = bc.Blocks[len(bc.Blocks)-1].Hash
	}

	newBlock := mineBlock(data, prevHash, bc.Difficulty)
	bc.Blocks = append(bc.Blocks, newBlock)
}

func main() {
	blockchain := Blockchain{Difficulty: 4}

	values := []int{91911, 90954, 95590, 97390, 96578, 97211, 95090}

	for _, value := range values {
		blockchain.addBlock(value)
	}

	fmt.Println("\nBlockchain:")
	for i, block := range blockchain.Blocks {
		fmt.Printf("Block %d: %+v\n", i, block)
	}
}
