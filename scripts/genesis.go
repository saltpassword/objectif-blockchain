package scripts

import (
    "time"
    "objectives-go/types"
)

func CreateGenesisBlock() types.Block {
    return types.Block{
        Index:     0,
        Timestamp: time.Now().String(),
        Data: types.DataBlock{
            ShortDesc: "Start of blockchain project",
            WhyDesc:   "To achieve personal objectives using blockchain technology.",
            GainDesc:  "Learning, deployment, visualization, and tool exploration.",
            EndDate:   "N/A",
            WhatsNext: "Continue adding new blocks with updated objectives.",
        },
        PrevHash: "",
        Hash:     "",
    }
}
