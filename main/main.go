package main

import (
    "fmt"
    "objectives-go/scripts"
    "objectives-go/types"
)

func main() {
    genesisBlock := scripts.CreateGenesisBlock()

    genesisBlock.Hash = scripts.CalculateHash(genesisBlock)
    fmt.Println("Genesis Block:")
    scripts.PrintBlock(genesisBlock)

    newBlockData := types.DataBlock{
        ShortDesc: "",
        WhyDesc:   "",
        GainDesc:  "",
        EndDate:   "",
        WhatsNext: "",
    }
    newBlock := scripts.GenerateBlock(genesisBlock, newBlockData)
    fmt.Println("\nNew Block:")
    scripts.PrintBlock(newBlock)
}
