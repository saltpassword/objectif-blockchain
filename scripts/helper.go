package scripts

import (
    "crypto/sha256"
    "encoding/hex"
    "objectives-go/types"
	"time"
	"fmt"
)

func CalculateHash(block types.Block) string {
    record := string(block.Index) + block.Timestamp + block.Data.ShortDesc + block.PrevHash
    h := sha256.New()
    h.Write([]byte(record))
    hashed := h.Sum(nil)
    return hex.EncodeToString(hashed)
}

func IsBlockValid(newBlock, oldBlock types.Block) bool {
    if oldBlock.Index+1 != newBlock.Index {
        return false
    }

    if oldBlock.Hash != newBlock.PrevHash {
        return false
    }

    if CalculateHash(newBlock) != newBlock.Hash {
        return false
    }

    return true
}

func GenerateBlock(oldBlock types.Block, data types.DataBlock) types.Block {
    var newBlock types.Block

    newBlock.Index = oldBlock.Index + 1
    newBlock.Timestamp = time.Now().String()
    newBlock.Data = data
    newBlock.PrevHash = oldBlock.Hash
    newBlock.Hash = CalculateHash(newBlock)

    return newBlock
}

func PrintBlock(block types.Block) {
    fmt.Println("Index:", block.Index)
    fmt.Println("Timestamp:", block.Timestamp)
    fmt.Println("Short Description:", block.Data.ShortDesc)
    fmt.Println("Why Description:", block.Data.WhyDesc)
    fmt.Println("Gain Description:", block.Data.GainDesc)
    fmt.Println("End Date:", block.Data.EndDate)
    fmt.Println("Whats Next:", block.Data.WhatsNext)
    fmt.Println("Previous Hash:", block.PrevHash)
    fmt.Println("Hash:", block.Hash)
    fmt.Println()
}