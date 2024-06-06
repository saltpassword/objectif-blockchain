package types

type Block struct {
    Index     int
    Timestamp string
    Data      DataBlock
    PrevHash  string
    Hash      string
}
